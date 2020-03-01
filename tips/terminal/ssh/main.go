package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	port int
	user string
)

var rootCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh [hostname] [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		if err := sshTerminal(&Config{
			Host: args[0],
			Port: port,
			User: user,
		}); err != nil {
			fmt.Printf("Failed sshTerminal: err=%s", err.Error())
		}
		fmt.Println("Completed")
	},
}

func main() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 22, "port")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", os.Getenv("USER"), "port")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed Execute: err=%s", err.Error())
	}
}

type Config struct {
	Host string
	Port int
	User string
}

func sshTerminal(conf *Config) (err error) {
	// Create ssh.ClientConfig
	sock, tmpErr := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if tmpErr != nil {
		err = fmt.Errorf("Failed Dial $SSH_AUTH_SOCK: err=%s", tmpErr.Error())
		return
	}
	agentsock := agent.NewClient(sock)
	signers, tmpErr := agentsock.Signers()
	if tmpErr != nil {
		err = fmt.Errorf("Failed Signers: %s", tmpErr.Error())
		return
	}
	sshConfig := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signers...),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH connect
	client, tmpErr := ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port), sshConfig)
	if tmpErr != nil {
		err = fmt.Errorf("Failed ssh.Dial: host=%s, port=%d, err=%s", conf.Host, conf.Port, tmpErr.Error())
		return
	}
	session, tmpErr := client.NewSession()
	if tmpErr != nil {
		err = fmt.Errorf("Failed NewSession: err=%s", tmpErr.Error())
		return
	}
	defer session.Close()

	// キー入力を接続先が認識できる形式に変換する
	fd := int(os.Stdin.Fd())
	state, tmpErr := terminal.MakeRaw(fd)
	if tmpErr != nil {
		err = fmt.Errorf("Failed MakeRaw: err=%s", tmpErr.Error())
		return
	}
	defer terminal.Restore(fd, state)

	// ターミナルサイズの取得
	w, h, tmpErr := terminal.GetSize(fd)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetSize: err=%s", tmpErr.Error())
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	tmpErr = session.RequestPty("xterm", h, w, modes)
	if tmpErr != nil {
		err = fmt.Errorf("Failed RequestPty: err=%s", tmpErr.Error())
		return
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	tmpErr = session.Shell()
	if tmpErr != nil {
		err = fmt.Errorf("Failed Shell: %s", tmpErr.Error())
		return
	}

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGWINCH)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGWINCH:
				fd := int(os.Stdout.Fd())
				w, h, _ = terminal.GetSize(fd)
				session.WindowChange(h, w)
			}
		}
	}()

	tmpErr = session.Wait()
	if tmpErr != nil {
		err = fmt.Errorf("Failed session.Wait: err=%s", tmpErr.Error())
		return
	}
	return
}
