package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	yaml "gopkg.in/yaml.v2"

	"github.com/exitshell/konnect/proxy"
)

// Konnect is a collection of SSHProxy objects.
type Konnect struct {
	Hosts         map[string]*proxy.SSHProxy `yaml:"hosts"`
	ProxyChan     chan bool                  `yaml:"-"`
	CompletedChan chan bool                  `yaml:"-"`
}

// Get an SSHProxy object by name.
func (k *Konnect) Get(name string) (*proxy.SSHProxy, error) {
	proxy, ok := k.Hosts[name]
	// Return error if SSHProxy rule is not found.
	if !ok {
		errMsg := fmt.Sprintf("Host '%v' not found", name)
		return proxy, errors.New(errMsg)
	}
	return proxy, nil
}

// GetHosts - Get host names in sorted order (asc).
func (k *Konnect) GetHosts() []string {
	names := []string{}
	for host := range k.Hosts {
		names = append(names, host)
	}
	sort.Strings(names)
	return names
}

// CheckHosts - Ensure that the given host names exist.
func (k *Konnect) CheckHosts(hosts []string) error {
	// If a given host does not exist
	// in Konnect.Hosts, then throw an error.
	for _, host := range hosts {
		if _, ok := k.Hosts[host]; ok != true {
			return fmt.Errorf("Undefined host %v", host)
		}
	}
	return nil
}

// LoadFromFile - Load and validate SSHProxy objects from a yaml config file.
func (k *Konnect) LoadFromFile(filename string) error {
	// Read config file.
	byteStr, err := ioutil.ReadFile(filename)
	if err != nil {
		errMsg := fmt.Sprintf("Config read error %v", err)
		return errors.New(errMsg)
	}

	// Populate a Konnect struct from a config file.
	err = yaml.Unmarshal(byteStr, k)
	if err != nil {
		errMsg := fmt.Sprintf("Config parse error %v", err)
		return errors.New(errMsg)
	}

	// Validate each SSHProxy in Konnect.
	for name, proxy := range k.Hosts {
		proxy.Filename = filename
		proxy.Name = name
		if err := proxy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// TestConnection - Test proxy ssh connection.
func (k *Konnect) TestConnection(host string) {
	proxy := k.Hosts[host]
	proxy.TestConnection()
	k.ProxyChan <- true
}

// Status - Check the status of one or more hosts.
func (k *Konnect) Status(hosts []string) {
	// For each specified host, launch a goroutine
	// to test the ssh connection.
	for _, host := range hosts {
		go k.TestConnection(host)
	}

	// This goroutine blocks until all specified
	// host connections have been tested.
	go func() {
		for i := 0; i < len(hosts); i++ {
			<-k.ProxyChan
		}
		k.CompletedChan <- true
	}()

	// Block until above goroutine completes.
	<-k.CompletedChan
	// Done

	// Print the status of all hosts.
	for _, host := range hosts {
		proxy := k.Hosts[host]
		proxy.PrintStatus()
	}
}
