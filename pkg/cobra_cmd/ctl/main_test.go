package ctl

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func TestMain(t *testing.T) {
	os.Args = []string{"resource", "compute", "create"}
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
}
