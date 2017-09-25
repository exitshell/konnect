package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tunedmystic/konnect/engine"
)

// ListCmd - List all hosts from config file.
var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all hosts",
	Long:  "List all hosts",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename := resolveFilename(cmd)

		// List all hosts.
		engine.Init(filename).List()
	},
}
