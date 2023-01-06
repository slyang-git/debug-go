package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "debug",
		Short: "debug go version",
		Long:  "A fast remote debugging environment configurator.",
		// Run: func(cmd *cobra.Command, args []string){
		// 	// do something
		// 	fmt.Println("curl -sSfL https://raw.githubusercontent.com/slyang-git/debug/main/install.sh | bash")
		// },
	}
)

func Execute() error {
	return rootCmd.Execute()
}


func init() {
	// do some initialize task here
}