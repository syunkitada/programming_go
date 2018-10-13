package compute

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "manage compute",
	Long: `manage compute
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compute remove")
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
