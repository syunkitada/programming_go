package image

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "manage image",
	Long: `manage image
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image remove")
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
