package image

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "manage image",
	Long: `manage image
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image create")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
