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
		hostList := ""
		for _, host := range konnect.GetHosts() {
			hostInfo := konnect.Hosts[host].Info()
			hostList += fmt.Sprintln(hostInfo)
		}
		fmt.Println(hostList)
	},
}
