package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	socketType    string
	inputLogFile  string
	outputLogFile string
)

var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "console [socket path]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		socketPath := args[0]

		if err := StartConsole(&Config{
			SocketType:    socketType,
			SocketPath:    socketPath,
			InputLogFile:  inputLogFile,
			OutputLogFile: outputLogFile,
		}); err != nil {
			fmt.Printf("Failed startConsole: err=%s\n", err.Error())
		}
		fmt.Println("Completed")
	},
}

func main() {
	rootCmd.PersistentFlags().StringVarP(
		&socketType, "type", "t", "unix", "socket type(unix)")
	rootCmd.PersistentFlags().StringVarP(
		&inputLogFile, "input", "i", "/tmp/console.input.log", "input log file path")
	rootCmd.PersistentFlags().StringVarP(
		&outputLogFile, "output", "o", "/tmp/console.output.log", "output log file path")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed Execute: err=%s", err.Error())
	}
}

type Config struct {
	SocketType    string
	SocketPath    string
	InputLogFile  string
	OutputLogFile string
}

func StartConsole(conf *Config) (err error) {
	var c net.Conn
	c, err = net.DialTimeout(conf.SocketType, conf.SocketPath, time.Duration(3*time.Second))
	if err != nil {
		return
	}
	defer c.Close()

	inputLogfile, tmpErr := os.OpenFile(conf.InputLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if tmpErr != nil {
		err = fmt.Errorf("Failed OpenFile")
		return
	}
	defer inputLogfile.Close()

	outputLogfile, tmpErr := os.OpenFile(conf.OutputLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if tmpErr != nil {
		err = fmt.Errorf("Failed OpenFile")
		return
	}
	defer outputLogfile.Close()

	// enter raw mode
	fd := int(os.Stdin.Fd())
	state, tmpErr := terminal.MakeRaw(fd)
	if tmpErr != nil {
		err = fmt.Errorf("Failed MakeRaw: err=%s", tmpErr.Error())
		return
	}

	chMutex := sync.Mutex{}
	isDone := false
	doneCh := make(chan bool, 2)
	readCh := make(chan string, 10)
	sigCh := make(chan os.Signal, 1)

	defer func() {
		chMutex.Lock()
		isDone = true
		close(doneCh)
		close(readCh)
		close(sigCh)
		chMutex.Unlock()
	}()

	// killされたときにシグナルを受け取る
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGKILL)

	isPreExitKey := false
	go func() {
		stdin := bufio.NewReader(os.Stdin)
		for {
			ch, tmpErr := stdin.ReadByte()
			if tmpErr == io.EOF {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed ReadByte: EOF")
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}

			// logging for debug
			fmt.Fprintf(inputLogfile, "%s:%d\n", string(ch), ch)

			// ^]. でexitする
			if ch == 29 { // 29 == ^]
				isPreExitKey = true
			} else if isPreExitKey && ch == 46 { // 46 == .
				chMutex.Lock()
				if !isDone {
					doneCh <- true
				}
				chMutex.Unlock()
				return
			} else {
				isPreExitKey = false
			}

			_, tmpErr = c.Write([]byte(string(ch)))
			if tmpErr != nil {
				err = fmt.Errorf("Failed Write")
				chMutex.Lock()
				if !isDone {
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, tmpErr := c.Read(buf[:])
			if tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed Read: err=%s", tmpErr.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			fmt.Fprint(outputLogfile, string(buf[0:n]))
			readCh <- string(buf[0:n])
		}
	}()

	// 逐次出力せずに、バッファしてから出力する
	// 10msec 出力が途切れたら(timeoutしたら)、まとめて出力する
	var strs []string
	timeout := time.Duration(60 * time.Second)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		select {
		case ch := <-sigCh:
			cancel()
			terminal.Restore(fd, state)
			log.Printf("\nExit by %s\n", ch.String())
			return
		case <-doneCh:
			cancel()
			terminal.Restore(fd, state)
			log.Printf("\nExit by doneCh\n")
			return
		case str := <-readCh:
			cancel()
			strs = append(strs, str)
			timeout = time.Duration(10 * time.Millisecond)
		case <-ctx.Done():
			cancel()
			fmt.Print(strings.Join(strs, ""))
			strs = []string{}
			timeout = time.Duration(60 * time.Second)
		}
	}

	return
}
