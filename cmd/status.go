package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/exitshell/konnect/engine"
	"github.com/spf13/cobra"
)

var allHosts bool

// StatusCmd - Check the status of one or more hosts.
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of one or more hosts",
	Long:  "Check the status of one or more hosts",
	Run: func(cmd *cobra.Command, args []string) {
		// Resolve filename from flags.
		filename := resolveFilename(cmd)
		// Init engine.
		konnect := engine.Init(filename)

		hosts := args

		// If `allHosts` is specified, then use
		// all hosts in the Konnect engine.
		if allHosts == true {
			hosts = konnect.GetHosts()
		}

		if allHosts == true && len(args) > 0 {
			log.Fatal("Cannot use --all with specific hosts")
		}

		// Check that at least one host was specified.
		if allHosts == false && len(args) == 0 {
			log.Fatal("Please specify one or more hosts")
		}

		// Remove duplicate host names.
		hosts = removeDuplicates(hosts)

		// Validate hosts.
		konnect.CheckHosts(hosts)

		// Check status of the resolved hosts.
		fmt.Printf("Testing connections for %v\n\n", strings.Join(hosts, ", "))
		konnect.Status(hosts)
	},
}

func init() {
	// Test connections for all hosts.
	StatusCmd.Flags().BoolVarP(&allHosts, "all", "a", false, "Test connections for all hosts")
}
