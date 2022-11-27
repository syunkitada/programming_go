package compute

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "compute",
	Short: "manage compute",
	Long: `manage compute
                This is sample description1.
                This is sample description2.`,
}
