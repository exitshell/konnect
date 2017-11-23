package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// AppVersion info.
var AppVersion string

// AppBuild info.
var AppBuild string

// AppDate info.
var AppDate string

var config string
var interactive bool
var version bool

// RootCmd - Entry point to the application.
var RootCmd = &cobra.Command{
	Use:   "konnect",
	Short: "Connect to SSH hosts.",
	Long:  "Konnect is a tool to define and connect to SSH hosts.",
	Run: func(cmd *cobra.Command, args []string) {
		// Perform interactive connect prompt.
		if err := interactivePrompt(cmd); err != nil {
			if err == errConfigNotFound {
				cmd.Help()
				fmt.Println()
			}
			handleErr(err)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(AppVersion)
			os.Exit(0)
		}
	},
}

func init() {
	// Config filename.
	RootCmd.PersistentFlags().StringVarP(&config, "filename", "f", "", "Specify config file")
	// Show version information.
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "View version information")
}

// AddCommands - Connects subcommands to the RootCmd.
func AddCommands() {
	ConnectCmd.AddCommand(TaskCmd)
	RootCmd.AddCommand(ArgsCmd)
	RootCmd.AddCommand(ConnectCmd)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(EditCmd)
	RootCmd.AddCommand(StatusCmd)
	RootCmd.AddCommand(VersionCmd)
}

// Execute - runs the RootCmd.
func Execute() error {
	AddCommands()
	return RootCmd.Execute()
}
