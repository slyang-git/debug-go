package commands

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "debug",
		Short: "debug go version",
		Long:  "A fast remote debugging environment configurator.",
		Run: func(cmd *cobra.Command, args []string){
			// do something
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
