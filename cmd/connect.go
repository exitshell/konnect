package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var hostName string
var taskName string

// ConnectCmd - Connect to a host.
var ConnectCmd = &cobra.Command{
	Use:   "to",
	Short: "Connect to a host",
	Long:  "Connect to a host",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[connectcmd] %v (%T)\n", args, args)

		if len(args) == 0 {
			log.Fatal("Please specify a host")
		}

		// Set hostname.
		hostName = args[0]

		if len(args) == 1 {
			// Connect to host.
			err := connectToHost(cmd, hostName, "")
			handleErr(err)
			os.Exit(0)
		}

		// Find the subcommand.
		subCmd, subArgs, err := cmd.Find(args[1:])
		handleErr(err)
		// If the subcommand is the same as the original command,
		// then no new subcommand was found. In that case, exit
		// with an error.
		if subCmd.Use == cmd.Use {
			log.Fatalf("Cannot parse subcommand %v\n", subArgs)
		} else {
			subCmd.Run(cmd, subArgs)
		}
	},
}

// TaskCmd - Run a task on a host.
var TaskCmd = &cobra.Command{
	Use:   "and",
	Short: "Run a task on a host",
	Long:  "Run a task on a host",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[taskcmd] %v (%T)\n", args, args)

		if len(args) == 0 {
			log.Fatal("Please specify a task")
		}

		// Set task name.
		taskName = args[0]

		// Connect to host and run a command.
		err := connectToHost(cmd, hostName, taskName)
		handleErr(err)
	},
}
