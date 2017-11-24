package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/exitshell/konnect/engine"
	"github.com/spf13/cobra"
)

// An interactivePrompt to connect to hosts.
func interactivePrompt(cmd *cobra.Command) error {
	// Resolve filename from flags.
	filename, err := resolveFilename(cmd)
	if err != nil {
		return err
	}
	fmt.Println(filename)

	// Init engine.
	konnect, err := engine.Init(filename)
	if err != nil {
		return err
	}

	// Get host names.
	hosts := konnect.GetHostNames()

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
		return errNoHostSelected
	}

	// Get proxy.
	proxy, err := konnect.GetHost(answer.Hostname)
	if err != nil {
		return err
	}

	// Connect to host.
	return proxy.Connect()
}
