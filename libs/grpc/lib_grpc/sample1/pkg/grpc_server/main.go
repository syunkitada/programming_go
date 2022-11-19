package grpc_server

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "grpc-server",
	Long: `grpc-server
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewSimpleServer(&config.Conf)
		if err := server.Serv(); err != nil {
			glog.Fatal(err)
		}
	},
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	config.InitFlags(rootCmd)
}
