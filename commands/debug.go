package commands

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "debug",
		Short: "debug go version",
		Long:  "A fast remote debugging environment configurator.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
