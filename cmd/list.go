package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tunedmystic/konnect/engine"
)

// ListCmd - List all hosts from config file.
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all hosts",
	Long:  "List all hosts",
	Run: func(cmd *cobra.Command, args []string) {
		// Get config filename from flags.
		filename, _ := cmd.Flags().GetString("filename")

		// List all hosts.
		engine.Init(filename).List()
	},
}
