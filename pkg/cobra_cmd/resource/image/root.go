package image

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "image",
	Short: "manage image",
	Long: `manage image
                This is sample description1.
                This is sample description2.`,
}
