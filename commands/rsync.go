package commands

import (
	"github.com/spf13/cobra"
)

var (
	rsyncCmd = &cobra.Command{
		Use: "rsync",
	}
)

func init() {
	rootCmd.AddCommand(rsyncCmd)
}
