package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/exitshell/konnect/engine"
	"github.com/exitshell/konnect/proxy"
	"github.com/exitshell/konnect/task"
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
	return "", errConfigNotFound
}

func makeDefaultConfig(filename string) error {
	// Make default proxylist.
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
	// Make default tasklist.
	taskList := map[string]*task.SSHTask{
		"ping": &task.SSHTask{
			Command: "echo ping",
		},
		"tailsys": &task.SSHTask{
			Command: "tail -f -n 100 /var/log/syslog",
		},
	}

	// Init engine, and assign structs.
	konnect := engine.New()
	konnect.Hosts = proxyList
	konnect.Tasks = taskList

	// Marshal the konnect hosts.
	hostsByteSlice, err := konnect.MarshalHosts()
	if err != nil {
		return err
	}

	// Marshal the konnect tasks.
	tasksByteSlice, err := konnect.MarshalTasks()
	if err != nil {
		return err
	}

	// Build the byte slice to export.
	data := []byte{}
	data = append(data, []byte(`# Configuration file for konnect.`)...)
	data = append(data, 10, 10)

	// Append hosts.
	data = append(data, []byte(`# Define your hosts here.`)...)
	data = append(data, 10)
	data = append(data, hostsByteSlice...)
	data = append(data, 10, 10)

	// Append tasks.
	data = append(data, []byte(`# Define your tasks here.`)...)
	data = append(data, 10)
	data = append(data, tasksByteSlice...)

	// Write byte slice to file.
	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	fmt.Println("Created configuration file at:")
	c := color.New(color.FgCyan, color.Bold)
	c.Printf("%v\n", filename)

	return nil
}

func connectToHost(cmd *cobra.Command, hostName, taskName string) error {
	// Resolve filename from flags.
	filename, err := resolveFilename(cmd)
	handleErr(err)

	// Init engine.
	konnect, err := engine.Init(filename)
	handleErr(err)

	// Get host.
	proxy, err := konnect.GetHost(hostName)
	handleErr(err)

	// Get task if specified.
	if taskName != "" {
		task, err := konnect.GetTask(taskName)
		handleErr(err)

		// Add task command to proxy.
		proxy.ExtraArgs = task.Command
	}

	// Connect to host.
	if err := proxy.Connect(); err != nil {
		log.Fatal(err)
	}

	return nil
}
