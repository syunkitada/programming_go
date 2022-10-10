package grpc_client

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/grpc-sample/pkg/config"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status",
	Run: func(cmd *cobra.Command, args []string) {
		client := NewSimpleClient(&config.Conf)
		reply, err := client.Status()
		glog.Info(reply)
		glog.Info(err)
	},
}

var getLogsCmd = &cobra.Command{
	Use:   "get-logs",
	Short: "get-logs",
	Run: func(cmd *cobra.Command, args []string) {
		client := NewSimpleClient(&config.Conf)
		client.GetLogs()
	},
}

var reportLogsCmd = &cobra.Command{
	Use:   "report-logs",
	Short: "report-logs",
	Run: func(cmd *cobra.Command, args []string) {
		client := NewSimpleClient(&config.Conf)
		client.ReportLogs()
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat",
	Run: func(cmd *cobra.Command, args []string) {
		client := NewSimpleClient(&config.Conf)
		client.Chat()
	},
}
