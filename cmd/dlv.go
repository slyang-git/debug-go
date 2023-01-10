// Real command behind debug dlv command:
// 	nohup dlv attach ${ProcessID} --listen=:${Port} --headless --accept-multiclient --continue --api-version=2 --log --log-output=rpc,dap,debugger >> /tmp/dlv.log 2>&1 &
// so we need give the process ID to be debuged on and listen port number. The port number can set a default one, like 8017.
// The process ID can be supplied in three ways:
// 1) supply directly in debug dlv [processID] [listen port], e.g.: debug dlv -i 1234 -p 8017
// 2) supply by read a config file. (process ID cant't be read as a fix value cause it aways changing during debugging sessions);
// 3) supply a process ID table, users can select the right one;
//
// When a user only run `debug dlv`, it will read a config file firstly. If there is no config file, it will promote a error msg ask user to supply a process ID.

package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ProcessID  string
	ListenPort string

	ProcessTableHeader []string
	ProcessRecords     []map[string]string

	dlvCmd = &cobra.Command{
		Use:   "dlv",
		Short: "Start a dlv process for debugging",
		Long:  `hello world hello world`,
		Run: func(cmd *cobra.Command, args []string) {
			// do something here
			// fmt.Println(args)
			ProcessID, _ = cmd.Flags().GetString("pid")
			ListenPort, _ = cmd.Flags().GetString("port")
			fmt.Println(ProcessID, ListenPort)

			command := exec.Command("ps", "-el")

			stdout, err := command.StdoutPipe()
			if err != nil {
				fmt.Println(err)
			}

			command.Start()

			buf := bufio.NewReader(stdout)

			line, _, _ := buf.ReadLine()
			for _, v := range strings.Fields(string(line)) {
				fmt.Println(v)
				ProcessTableHeader = append(ProcessTableHeader, v)
			}
			fmt.Println(ProcessTableHeader)

			for len(line) > 0 {
				line, _, _ = buf.ReadLine()

				tmp := make(map[string]string)
				for i, v := range strings.Fields(string(line)) {
					tmp[ProcessTableHeader[i]] = v
				}
				ProcessRecords = append(ProcessRecords, tmp)
				// fmt.Println(string(line))
			}

			// fmt.Println(ProcessRecords)
			for _, v := range ProcessRecords {
				for k, vv := range v {
					fmt.Println(k, vv)
				}
            }
			

		},
	}
)

func init() {
	rootCmd.AddCommand(dlvCmd)

	dlvCmd.Flags().StringVarP(&ProcessID, "pid", "i", "", "The process ID you gonna debugging")
	dlvCmd.Flags().StringVarP(&ListenPort, "port", "p", "8018", "The listen port of dlv process")
}
