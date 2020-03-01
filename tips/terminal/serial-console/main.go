package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "console [socket path]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		if err := startConsole(args[0]); err != nil {
			fmt.Printf("Failed startConsole: err=%s", err.Error())
		}
		fmt.Println("Completed")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed Execute: err=%s", err.Error())
	}
}

func startConsole(socketPath string) (err error) {
	// https://github.com/golang/crypto/blob/master/ssh/terminal/terminal.go

	// goapp@ubuntu:/$ stty -a
	// stty -a
	// speed 9600 baud; rows 24; columns 80; line = 0;
	// intr = ^C; quit = ^\; erase = ^?; kill = ^U; eof = ^D; eol = <undef>;
	// eol2 = <undef>; swtch = <undef>; start = ^Q; stop = ^S; susp = ^Z; rprnt = ^R;
	// werase = ^W; lnext = ^V; discard = ^O; min = 1; time = 0;
	// -parenb -parodd -cmspar cs8 hupcl -cstopb cread clocal -crtscts
	// -ignbrk -brkint -ignpar -parmrk -inpck -istrip -inlcr -igncr icrnl ixon ixoff
	// -iuclc -ixany -imaxbel iutf8
	// opost -olcuc -ocrnl onlcr -onocr -onlret -ofill -ofdel nl0 cr0 tab0 bs0 vt0 ff0
	// isig icanon -iexten echo echoe echok -echonl -noflsh -xcase -tostop -echoprt
	// echoctl echoke -flusho -extproc

	var c net.Conn
	c, err = net.Dial("unix", socketPath)
	if err != nil {
		return
	}
	defer c.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, tmpErr := c.Read(buf[:])
			if tmpErr != nil {
				log.Printf("Failed Read: err=%s", tmpErr.Error())
				return
			}

			fmt.Print(string(buf[0:n]))
		}
	}()

	stdin := bufio.NewReader(os.Stdin)
	for {
		ch, tmpErr := stdin.ReadByte()
		if tmpErr == io.EOF {
			fmt.Println("EOF")
			break
		}
		_, tmpErr = c.Write([]byte(string(ch)))
		if tmpErr != nil {
			err = fmt.Errorf("Failed Write")
			return
		}
	}
	return
}
