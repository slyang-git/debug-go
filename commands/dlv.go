package commands

import (
	"github.com/spf13/cobra"
)

var (
	dlvCmd = &cobra.Command{
		Use: "dlv",
		Short: "hello",
		Long: "hello world hello world",
	}
)

func init() {
	rootCmd.AddCommand(dlvCmd)
}
