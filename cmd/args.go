package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tunedmystic/konnect/engine"
)

// ArgsCmd - Print the SSH command for a host.
var ArgsCmd = &cobra.Command{
	Use:   "args",
	Short: "Print the SSH command for a host",
	Long:  "Print the SSH command for a host",
	Run: func(cmd *cobra.Command, args []string) {
		// Get config filename from flags.
		filename, _ := cmd.Flags().GetString("filename")

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}

		// Print Host SSH command.
		engine.Init(filename).Args(args[0])
	},
}
