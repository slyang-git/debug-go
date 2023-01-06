package commands

import (
	"github.com/spf13/cobra"
)

var (
	rsyncCmd = &cobra.Command{
		Use: "rsync",
		Short: "rsync is for upload files",
		Long: "hello world hello world",
	}
)

func init() {
	rootCmd.AddCommand(rsyncCmd)
}
