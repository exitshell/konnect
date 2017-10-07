package cmd

import (
	"fmt"
	"log"

	"github.com/exitshell/konnect/engine"
	"github.com/spf13/cobra"
)

// ArgsCmd - Print the SSH command for a host.
var ArgsCmd = &cobra.Command{
	Use:   "args",
	Short: "Print the SSH command for a host",
	Long:  "Print the SSH command for a host",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename, err := resolveFilename(cmd)
		if err != nil {
			log.Fatal(err)
		}

		// Check that only one host was specified.
		if len(args) != 1 {
			log.Fatal("Please specify one host")
		}

		// Init engine.
		konnect, err := engine.Init(filename)
		if err != nil {
			log.Fatal(err)
		}

		// Print Host SSH command.
		hostArgs, err := konnect.Args(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hostArgs)
	},
}
