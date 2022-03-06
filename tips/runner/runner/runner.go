package runner

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const (
	StatusTimeout = 124
)

type Config struct {
	Timeout  int
	Interval int
	Cmd      string
	UseShell bool
}

type Runner struct {
	conf         *Config
	cmd          string
	cmdOptions   []string
	cmdTimeout   time.Duration
	killLimit    int
	killInterval time.Duration
	stopCh       chan chan bool
	isStarted    bool
}

func New(conf *Config) *Runner {
	cmds := strings.Fields(conf.Cmd)
	cmd := cmds[0]
	cmdOptions := cmds[1:]
	timeout := conf.Timeout
	if timeout == 0 {
		timeout = conf.Interval - 10
	}
	cmdTimeout := time.Duration(timeout) * time.Second
	return &Runner{
		conf:         conf,
		cmd:          cmd,
		cmdOptions:   cmdOptions,
		cmdTimeout:   cmdTimeout,
		killLimit:    2,
		killInterval: time.Duration(1) * time.Second,
		stopCh:       make(chan chan bool),
	}
}

func (self *Runner) Start() {
	fmt.Println("DEBUG Start")
	if !self.isStarted {
		go self.start()
		self.isStarted = true
	} else {
		log.Printf("Already Started")
	}
}

func (self *Runner) Stop() {
	log.Printf("stopping: %v", self.conf.Cmd)
	doneCh := make(chan bool)
	self.stopCh <- doneCh
	<-doneCh
	fmt.Println("end stop")
}

func (self *Runner) start() {
	interval := time.Duration(self.conf.Interval) * time.Second
	ticker := time.NewTicker(interval)
	log.Printf("start: %s", self.conf.Cmd)

	self.Run()
	for {
		select {
		case doneCh := <-self.stopCh:
			doneCh <- true
			log.Printf("done: %s", self.conf.Cmd)
			return
		case t := <-ticker.C:
			fmt.Println("tick at", t)
			self.Run()
		}
	}
	return
}

type Result struct {
	Cmd    string
	Err    error
	Output string
	Status int
}

func (self *Runner) MustKillProcess(pid int) {
	if pid == 0 {
		return
	}
	for i := 0; i < self.killLimit; i++ {
		process, err := GetProcess(pid)
		if err != nil {
			log.Fatalf("Unexpected Error: %s", err.Error())
		}
		if process == nil {
			return
		}

		if self.conf.UseShell && process.Cmds[2] != self.conf.Cmd {
			log.Fatalf("Unexpected Cmd Found: expectedCmd=%s, foundCmd=%v", self.cmd, process.Cmds)
		} else if process.Cmds[0] != self.cmd {
			log.Fatalf("Unexpected Cmd Found: expectedCmd=%s, foundCmd=%v", self.cmd, process.Cmds)
		}

		log.Printf("ExistsProcess will be killed: pid=%d, cmds=%v", pid, process.Cmds)
		if err = syscall.Kill(-pid, syscall.SIGKILL); err != nil {
			log.Fatalf("Unexpected Error: %s", err.Error())
			return
		}
		time.Sleep(self.killInterval)
	}

	process, err := GetProcess(pid)
	if err != nil {
		log.Fatalf("Unexpected Error: %s", err.Error())
	}
	if process != nil {
		if process.Cmds[0] != self.cmd {
			log.Fatalf("Failed KillProcess, and Unexpected Cmd Found: expectedCmd=%s, foundCmd=%v", self.cmd, process.Cmds)
		}
		log.Fatalf("Failed KillProcess: pid=%d, cmds=%v", pid, process.Cmds)
	}
}

func (self *Runner) Run() (result *Result, err error) {
	log.Printf("Run: %s", self.conf.Cmd)
	var pid int
	var output []byte
	var status int

	defer func() {
		self.MustKillProcess(pid)

		result = &Result{
			Cmd:    self.conf.Cmd,
			Output: string(output),
			Status: status,
		}
		if status == StatusTimeout {
			result.Output = "command timeout: " + result.Output
		}
		log.Printf("EndRun: %s", self.conf.Cmd)
		return
	}()

	var command *exec.Cmd
	if self.conf.UseShell {
		command = exec.Command("sh", "-c", self.conf.Cmd)
	} else {
		command = exec.Command(self.cmd, self.cmdOptions...)
	}
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var stdout io.ReadCloser
	if stdout, err = command.StdoutPipe(); err != nil {
		return
	}

	resultCh := make(chan bool)
	go func() {
		if err = command.Start(); err != nil {
			return
		}
		pid = command.Process.Pid
		if output, err = ioutil.ReadAll(stdout); err != nil {
			return
		}
		if err = command.Wait(); err != nil {
			return
		}
		status = command.ProcessState.ExitCode()
		resultCh <- true
	}()

	select {
	case <-resultCh:
		return
	case <-time.After(self.cmdTimeout):
		status = StatusTimeout
	}
	return
}
