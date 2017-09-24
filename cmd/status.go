package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tunedmystic/konnect/engine"
)

// StatusCmd - Check the status of one or more hosts.
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of one or more hosts",
	Long:  "Check the status of one or more hosts",
	Run: func(cmd *cobra.Command, args []string) {
		// Get config filename from flags.
		filename, _ := cmd.Flags().GetString("filename")

		// Check that only one host was specified.
		if len(args) == 0 {
			log.Fatal("Please specify one or more hosts")
		}

		// Remove duplicate host names.
		hosts := removeDuplicates(args)

		// Connect to host.
		engine.Init(filename).Status(hosts)
	},
}
