package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version bool
var config string

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
	},
}

func init() {
	// Create flags for RootCmd.
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "View version information")
	RootCmd.PersistentFlags().StringVarP(&config, "filename", "f", getDefaultConfig(), "Specify config file")
}

// AddCommands - Connects subcommands to the RootCmd.
func AddCommands() {
	RootCmd.AddCommand(ArgsCmd)
	RootCmd.AddCommand(ConnectCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(VersionCmd)
}

// Execute - runs the RootCmd.
func Execute() error {
	AddCommands()
	return RootCmd.Execute()
}
