package proxylist

import (
	"errors"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	ssh "github.com/tunedmystic/konnect/sshproxy"
)

// ProxyList is a collection of SSHProxy objects.
type ProxyList struct {
	Hosts map[string]*ssh.SSHProxy `yaml:"hosts"`
}

// Get an SSHProxy object by name.
func (pl *ProxyList) Get(name string) (*ssh.SSHProxy, error) {
	proxy, ok := pl.Hosts[name]
	// Return error if SSHProxy rule is not found.
	if !ok {
		errMsg := fmt.Sprintf("[config] SSH Rule '%v' not found", name)
		return proxy, errors.New(errMsg)
	}
	return proxy, nil
}

// LoadFromFile - Load and validate SSHProxy objects from a yaml config file.
func (pl *ProxyList) LoadFromFile(filename string) error {
	// Read config file.
	byteStr, err := ioutil.ReadFile(filename)
	if err != nil {
		errMsg := fmt.Sprintf("[config] Read config file %v", err)
		return errors.New(errMsg)
	}

	// Inflate ProxyList from config file data.
	err = yaml.Unmarshal(byteStr, pl)
	if err != nil {
		errMsg := fmt.Sprintf("[config] Parse config file %v", err)
		return errors.New(errMsg)
	}

	// Validate each SSHProxy in ProxyList.
	for name, proxy := range pl.Hosts {
		proxy.Filename = filename
		proxy.Name = name
		if err := proxy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// New - Create a new ProxyList object.
func New() *ProxyList {
	return &ProxyList{
		Hosts: make(map[string]*ssh.SSHProxy),
	}
}
