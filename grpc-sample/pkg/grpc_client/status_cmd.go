package grpc_client

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status",
	Long: `status
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewSimpleClient(&config.Conf)
		reply, err := client.Status()
		glog.Info(reply)
		glog.Info(err)
	},
}
