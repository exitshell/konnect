package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/exitshell/konnect/engine"
	"github.com/exitshell/konnect/proxy"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// handleErr is a function that logs Fatal
// if the given error `err` is populated.
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// This function returns a slice of
// possible default config filenames.
func getDefaultConfigs() []string {
	return []string{
		"./konnect.yml",
		"../konnect.yml",
	}
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

// Get the config filename from cmd flags.
// Fallback to default filenames.
// Validate that the file exists.
func resolveFilename(cmd *cobra.Command) (string, error) {
	// Get config filename from flags.
	filename, _ := cmd.Flags().GetString("filename")
	filenames := []string{filename}

	// If filename is not specified, then get
	// a list of possible config filenames.
	if filename == "" {
		filenames = getDefaultConfigs()
	}

	for _, fName := range filenames {
		// Check if the filename exists.
		if _, err := os.Stat(fName); err == nil {
			// Filename was found. Immediately return.
			return fName, nil
		}
	}

	// At this point, none of the possible filenames
	// were found. Return an error.
	err := errors.New("Could not find a " +
		"konnect.yml configuration file.")
	return "", err
}

func makeDefaultConfig(filename string) error {
	// Make default proxylist and konnect engine.
	proxyList := map[string]*proxy.SSHProxy{
		"app": &proxy.SSHProxy{
			User: "root",
			Host: "127.0.0.1",
			Port: 22,
			Key:  "/home/app/key",
		},
		"database": &proxy.SSHProxy{
			User: "admin",
			Host: "192.168.99.100",
			Port: 89,
			Key:  "~/.ssh/id_rsa",
		},
	}
	konnect := engine.New()
	konnect.Hosts = proxyList

	// Marshal konnect struct to a byte slice.
	byteSlice, err := yaml.Marshal(konnect)
	if err != nil {
		return err
	}

	// Make config header.
	header := []byte{}
	header = append(header, []byte(`# Autogenerated configuration file.`)...)
	header = append(header, 10)
	header = append(header, []byte(`# Define your hosts here and connect to them with konnect cli.`)...)
	header = append(header, 10, 10)

	// Make config body.
	data := []byte{}
	data = append(data, header...)
	data = append(data, byteSlice...)

	// Write byte slice to file.
	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	fmt.Println("Created configuration file at:")
	c := color.New(color.FgCyan, color.Bold)
	c.Printf("%v\n", filename)

	return nil
}
