package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/cobra-sample/get"
)

var rootCmd = &cobra.Command{
	Use: "cobra-sample",
}

func init() {
	rootCmd.AddCommand(get.GetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
