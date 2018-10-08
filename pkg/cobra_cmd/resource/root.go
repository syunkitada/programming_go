package resource

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-sample/pkg/cobra_cmd/resource/compute"
	"github.com/syunkitada/go-sample/pkg/cobra_cmd/resource/image"
)

var RootCmd = &cobra.Command{
	Use:   "resource",
	Short: "manage resource",
	Long: `manage resource
                This is sample description1.
                This is sample description2.`,
}

func init() {
	RootCmd.AddCommand(compute.RootCmd)
	RootCmd.AddCommand(image.RootCmd)
}
