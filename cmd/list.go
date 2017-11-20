package cmd

import (
	"fmt"
	"log"

	"github.com/exitshell/konnect/engine"
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
		fmt.Printf("-- Hosts --\n\n")
		hostList := ""
		for _, hostName := range konnect.GetHostNames() {
			hostInfo := konnect.Hosts[hostName].Info()
			hostList += fmt.Sprintln(hostInfo)
		}
		fmt.Println(hostList)

		// Show info for all tasks.
		fmt.Printf("\n-- Tasks --\n\n")
		taskList := ""
		for _, taskName := range konnect.GetTaskNames() {
			taskInfo := konnect.Tasks[taskName].Info()
			taskList += fmt.Sprintln(taskInfo)
		}
		fmt.Println(taskList)
	},
}
