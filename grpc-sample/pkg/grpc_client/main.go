package grpc_client

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(statusCmd)
}
