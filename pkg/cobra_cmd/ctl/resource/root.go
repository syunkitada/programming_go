package resource

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/go-samples/pkg/cobra_cmd/ctl/resource/compute"
	"github.com/syunkitada/go-samples/pkg/cobra_cmd/ctl/resource/image"
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
