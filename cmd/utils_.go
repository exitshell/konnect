package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func getDefaultConfig() string {
	return "./konnect.yml"
}

func getVersion() string {
	version := "0.0.1"
	return version
}

// Remove duplicate elements from a string slice.
// https://goo.gl/ttDAg2
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for _, host := range elements {
		if encountered[host] == false {
			encountered[host] = true
			result = append(result, host)
		}
	}
	return result
}

// Resolve the config filename from cmd flags.
// Fallback to default filename.
// Validate that the file exists.
func resolveFilename(cmd *cobra.Command) string {
	// Get config filename from flags.
	filename, _ := cmd.Flags().GetString("filename")
	wasProvided := true

	// If filename is not specified, then set
	// to default config filename.
	if filename == "" {
		wasProvided = false
		filename = getDefaultConfig()
	}

	// Check if the filename exists.
	if _, err := os.Stat(filename); err != nil {
		// Could not find a config file.
		if wasProvided == false {
			err = errors.New("Could not find a " +
				"konnect.yml configuration file " +
				"in this directory")
		} else {
			// Could not find the config file that the user specified.
			err = fmt.Errorf("Config %v does not exist", filename)
		}
		log.Fatal(err)
	}

	return filename
}
