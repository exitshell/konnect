package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey"
	"github.com/exitshell/konnect/engine"
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
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(AppVersion)
			os.Exit(0)
		}

		if interactive {
			InteractivePrompt(cmd)
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

// InteractivePrompt to connect to hosts.
func InteractivePrompt(cmd *cobra.Command) {
	fmt.Println("Starting interactive prompt...")
	// Resolve filename from flags.
	filename, err := resolveFilename(cmd)
	handleErr(err)

	// Init engine.
	konnect, err := engine.Init(filename)
	handleErr(err)

	// Get host names.
	hosts := konnect.GetHosts()

	// Create survey.
	prompt := []*survey.Question{
		{
			Name:     "Hostname",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Connect to host:",
				Options: hosts,
			},
		},
	}

	// Create answer.
	answer := struct {
		Hostname string
	}{}

	// Show prompt.
	if err = survey.Ask(prompt, &answer); err != nil {
		log.Fatal("No host was selected")
	}

	// Get proxy.
	proxy, err := konnect.Get(answer.Hostname)
	handleErr(err)

	// Connect to host.
	if err := proxy.Connect(); err != nil {
		log.Fatal(err)
	}
}

// AddCommands - Connects subcommands to the RootCmd.
func AddCommands() {
	RootCmd.AddCommand(ArgsCmd)
	RootCmd.AddCommand(ConnectCmd)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(OpenCmd)
	RootCmd.AddCommand(StatusCmd)
	RootCmd.AddCommand(VersionCmd)
}

// Execute - runs the RootCmd.
func Execute() error {
	AddCommands()
	return RootCmd.Execute()
}
