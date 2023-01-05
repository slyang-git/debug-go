package commands

import (
	"github.com/spf13/cobra"
)

var (
	dlvCmd = &cobra.Command{
		Use: "dlv",
	}
)

func init() {
	rootCmd.AddCommand(dlvCmd)
}
