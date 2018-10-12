package image

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "manage image",
	Long: `manage image
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image list")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
