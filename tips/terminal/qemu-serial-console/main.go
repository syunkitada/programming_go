package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	c, err = net.Dial(conf.SocketType, conf.SocketPath)
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
	defer terminal.Restore(fd, state)

	go func() { // killされたときもterminalを明示的に戻す
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGKILL)
		select {
		case ch := <-sigCh:
			terminal.Restore(fd, state)
			log.Fatalf("\nExit by %s\n", ch.String())
		}

		terminal.Restore(fd, state)
		log.Fatalf("\nExit by Unknown\n")
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, tmpErr := c.Read(buf[:])
			if tmpErr != nil {
				// FIXME 正常終了できてない? 終了時にここに入ることがある
				log.Printf("Failed Read: err=%s", tmpErr.Error())
				return
			}

			// logging for debug
			fmt.Fprint(outputLogfile, string(buf[0:n]))

			fmt.Print(string(buf[0:n]))
		}
	}()

	done := make(chan bool)

	isPreExitKey := false
	go func() {
		stdin := bufio.NewReader(os.Stdin)
		for {
			ch, tmpErr := stdin.ReadByte()
			if tmpErr == io.EOF {
				err = fmt.Errorf("Failed ReadByte: EOF")
				done <- true
				return
			}

			// logging for debug
			fmt.Fprintf(inputLogfile, "%s:%d\n", string(ch), ch)

			// ^]. でexitする
			if ch == 29 { // 29 == ^]
				isPreExitKey = true
			} else if isPreExitKey && ch == 46 { // 46 == .
				done <- true
			} else {
				isPreExitKey = false
			}

			_, tmpErr = c.Write([]byte(string(ch)))
			if tmpErr != nil {
				err = fmt.Errorf("Failed Write")
				done <- true
				return
			}
		}
	}()

	select {
	case <-done:
		close(done)
	}
	return
}
