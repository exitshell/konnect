package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/exitshell/konnect/engine"
)

// ArgsCmd - Print the SSH command for a host.
var ArgsCmd = &cobra.Command{
	Use:   "args",
	Short: "Print the SSH command for a host",
	Long:  "Print the SSH command for a host",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename := resolveFilename(cmd)

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}

		// Print Host SSH command.
		engine.Init(filename).Args(args[0])
	},
}
