package main

import (
	"fmt"
	"os"

	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/gen/cli"
)

func main() {
	rootCmd, err := cli.MakeRootCmd()
	if err != nil {
		fmt.Println("Cmd construction error: ", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
