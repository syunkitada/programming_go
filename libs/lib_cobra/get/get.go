package get

import (
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource",
}

var hogeCmd = &cobra.Command{
	Use:   "hoge",
	Short: "Get hoge",
}

func init() {
	GetCmd.AddCommand(hogeCmd)
}
