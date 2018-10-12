package compute

import (
	"fmt"

	"github.com/spf13/cobra"
)

var output string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "manage compute",
	Long: `manage compute
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(output)
		fmt.Println("compute list")
	},
}

func init() {
	listCmd.Flags().StringVarP(&output, "output", "o", "json", "output format")
	RootCmd.AddCommand(listCmd)
}
