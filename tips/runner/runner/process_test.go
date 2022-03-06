package runner

import (
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetProcess(t *testing.T) {
	a := assert.New(t)

	var pid int
	command := exec.Command("sleep", "2")
	go func() {
		command.Start()
		pid = command.Process.Pid
		command.Wait()
	}()
	time.Sleep(1 * time.Second)

	{
		process, err := GetProcess(pid)
		a.Equal(err, nil)
		expected := Process{
			Pid:  process.Pid,
			Cmds: []string{"sleep", "2"},
		}
		a.Equal(expected, *process)
	}

	{
		process, err := GetProcess(-1)
		a.Equal(err, nil)
		var expectedProcess *Process
		a.Equal(expectedProcess, process)
	}
}
