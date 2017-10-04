package cmd

import (
	"log"

	"github.com/exitshell/konnect/engine"
	"github.com/spf13/cobra"
)

// ConnectCmd - Connect to a host.
var ConnectCmd = &cobra.Command{
	Use:   "to",
	Short: "Connect to a host",
	Long:  "Connect to a host",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename := resolveFilename(cmd)

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}

		// Init engine.
		konnect, err := engine.Init(filename)
		if err != nil {
			log.Fatal(err)
		}

		// Connect to host.
		if err := konnect.Connect(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}
