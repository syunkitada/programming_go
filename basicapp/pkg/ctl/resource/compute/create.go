package compute

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/basicapp/pkg/config"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "manage compute",
	Long: `manage compute
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.Info(config.Conf)
		glog.Info("DEBUG")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
