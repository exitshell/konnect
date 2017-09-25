package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var config string
var interactive bool
var version bool

// RootCmd - Entry point to the application.
var RootCmd = &cobra.Command{
	Use:   "konnect",
	Short: "Connect to SSH hosts.",
	Long:  "Define and connect to SSH hosts.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(getVersion())
			os.Exit(0)
		}

		if interactive {
			InteractivePrompt()
			os.Exit(0)
		}
	},
}

func init() {
	// Config filename.
	RootCmd.PersistentFlags().StringVarP(&config, "filename", "f", "", "Specify config file")
	// Show an iteractive prompt to connect to hosts.
	RootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Connect to a host interactively")
	// Show version information.
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "View version information")
}

// Iteractive prompt to connect to hosts.
func InteractivePrompt() {
	fmt.Println("Starting interactive prompt...")
}

// AddCommands - Connects subcommands to the RootCmd.
func AddCommands() {
	RootCmd.AddCommand(ArgsCmd)
	RootCmd.AddCommand(ConnectCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(StatusCmd)
	RootCmd.AddCommand(VersionCmd)
}

// Execute - runs the RootCmd.
func Execute() error {
	AddCommands()
	return RootCmd.Execute()
}
