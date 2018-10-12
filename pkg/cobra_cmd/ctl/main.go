package ctl

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
