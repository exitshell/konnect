package cmd

import (
	"fmt"
	"log"
	"strings"

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

		// Get args for host.
		argsStr := strings.Join(proxy.Args(), " ")
		fmt.Println(argsStr)
	},
}
