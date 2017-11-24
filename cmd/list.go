package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/exitshell/konnect/engine"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// ListCmd - List all hosts from config file.
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all hosts",
	Long:  "List all hosts",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename, err := resolveFilename(cmd)
		handleErr(err)

		// Check that only one host was specified.
		if len(args) != 0 {
			log.Fatal("The list subcommand does not take any arguments")
		}

		// Init engine.
		konnect, err := engine.Init(filename)
		handleErr(err)

		// Show info for all hosts.
		hostList := [][]string{}
		for _, hostName := range konnect.GetHostNames() {
			hostInfo := konnect.Hosts[hostName].Info()
			hostList = append(hostList, []string{
				hostName,
				hostInfo,
			})
		}
		hostTable := tablewriter.NewWriter(os.Stdout)
		hostTable.SetHeader([]string{"Hosts", ""})
		hostTable.AppendBulk(hostList)
		hostTable.Render()

		fmt.Printf("\n\n")

		// Show info for all tasks.
		taskList := [][]string{}
		for _, taskName := range konnect.GetTaskNames() {
			taskInfo := konnect.Tasks[taskName].Info()
			taskList = append(taskList, []string{
				taskName,
				taskInfo,
			})
		}
		taskTable := tablewriter.NewWriter(os.Stdout)
		taskTable.SetHeader([]string{"Tasks", ""})
		taskTable.AppendBulk(taskList)
		taskTable.Render()
	},
}
