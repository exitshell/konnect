package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	yaml "gopkg.in/yaml.v2"

	"github.com/tunedmystic/konnect/proxy"
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
		errMsg := fmt.Sprintf("[config] SSH Rule '%v' not found", name)
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
func (k *Konnect) CheckHosts(hosts []string) {
	// If a given host does not exist
	// in Konnect.Hosts, then throw an error.
	for _, host := range hosts {
		if _, ok := k.Hosts[host]; ok != true {
			log.Fatalf("undefined host %v", host)
		}
	}
}

// LoadFromFile - Load and validate SSHProxy objects from a yaml config file.
func (k *Konnect) LoadFromFile(filename string) error {
	// Read config file.
	byteStr, err := ioutil.ReadFile(filename)
	if err != nil {
		errMsg := fmt.Sprintf("[config] Read config file %v", err)
		return errors.New(errMsg)
	}

	// Populate a Konnect struct from a config file.
	err = yaml.Unmarshal(byteStr, k)
	if err != nil {
		errMsg := fmt.Sprintf("[config] Parse config file %v", err)
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

// New - Create a new Konnect object.
func New() *Konnect {
	return &Konnect{
		Hosts:         make(map[string]*proxy.SSHProxy),
		ProxyChan:     make(chan bool),
		CompletedChan: make(chan bool),
	}
}
