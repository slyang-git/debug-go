package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the debug's version number",
	Long:  "All software has versions. This is debug's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Debug's version is v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
