package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tunedmystic/konnect/engine"
)

// ConnectCmd - Connect to a host.
var ConnectCmd = &cobra.Command{
	Use:   "to",
	Short: "Connect to a host",
	Long:  "Connect to a host",
	Run: func(cmd *cobra.Command, args []string) {
		// Get config filename from flags.
		filename, _ := cmd.Flags().GetString("filename")

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}

		// Connect to host.
		engine.Init(filename).Connect(args[0])
	},
}
