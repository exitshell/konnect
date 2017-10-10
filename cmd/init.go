package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// InitCmd - Create configuration file.
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a konnect.yml configuration file at the given directory.",
	Long:  "Initialize a konnect.yml configuration file at the given directory.",
	Run: func(cmd *cobra.Command, args []string) {
		// Directory for the configuration file.
		// Default to the current directory.
		dir := "."

		// Check if too many args were given.
		if len(args) > 1 {
			log.Fatal("Too many args. Please specify a directory")
		}

		// If a diectory is specified, then use the given directory.
		if len(args) == 1 {
			dir = args[0]
		}

		// Get absolute path of the dir.
		dir, _ = filepath.Abs(dir)

		// Get the default name for the config file.
		localFilename := getDefaultConfig()

		// Join the resolved directory with
		// the default config filename.
		filename := filepath.Join(dir, localFilename)

		// If the file already exists, then return error and exit.
		if _, err := os.Stat(filename); err == nil {
			log.Fatalf("File %v already exists.\n", filename)
		}

		if err := makeDefaultConfig(filename); err != nil {
			log.Fatal(err)
		}
	},
}
