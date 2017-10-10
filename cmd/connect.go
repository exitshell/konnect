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
		filename, err := resolveFilename(cmd)
		handleErr(err)

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}
		host := args[0]

		// Init engine.
		konnect, err := engine.Init(filename)
		handleErr(err)

		// Get host.
		proxy, err := konnect.Get(host)
		handleErr(err)

		// Connect to host.
		if err := proxy.Connect(); err != nil {
			log.Fatal(err)
		}
	},
}
