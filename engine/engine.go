package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	yaml "gopkg.in/yaml.v2"

	"github.com/exitshell/konnect/proxy"
	"github.com/exitshell/konnect/task"
)

// Konnect is a collection of SSHProxy objects.
type Konnect struct {
	Global        *proxy.SSHProxy
	ProxyChan     chan bool
	CompletedChan chan bool
	Hosts         map[string]*proxy.SSHProxy
	Tasks         map[string]*task.SSHTask
}

// Get an SSHProxy object by name.
func (k *Konnect) GetHost(name string) (*proxy.SSHProxy, error) {
	proxy, ok := k.Hosts[name]
	// Return error if SSHProxy rule is not found.
	if !ok {
		errMsg := fmt.Sprintf("Host '%v' not found\n", name)
		return proxy, errors.New(errMsg)
	}
	return proxy, nil
}

// GetTask - Get an SSHTask object by name.
func (k *Konnect) GetTask(name string) (*task.SSHTask, error) {
	task, ok := k.Tasks[name]
	// Return error if SSHTask is not found.
	if !ok {
		errMsg := fmt.Sprintf("Task '%v' not found\n", name)
		return task, errors.New(errMsg)
	}
	return task, nil
}

// GetHostNames - Get host names in sorted order (asc).
func (k *Konnect) GetHostNames() []string {
	names := []string{}
	for host := range k.Hosts {
		names = append(names, host)
	}
	sort.Strings(names)
	return names
}

// GetTaskNames - Get task names in sorted order (asc).
func (k *Konnect) GetTaskNames() []string {
	names := []string{}
	for task := range k.Tasks {
		names = append(names, task)
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

// UnmarshalGlobal - Unmarshal global proxy config from a byte string.
func (k *Konnect) UnmarshalGlobal(byteStr []byte) error {
	// Make temporary struct to hold global data.
	var tempGlobal struct {
		Global *proxy.SSHProxy `yaml:"global"`
	}
	// Set the temp struct to an SSHProxy
	// with default global values.
	tempGlobal.Global = proxy.NewGlobal()

	// Unmarshal the byte string into the temp global struct.
	if err := yaml.Unmarshal(byteStr, &tempGlobal); err != nil {
		return err
	}
	// Assign to Konnect.
	k.Global = tempGlobal.Global

	return nil
}

// UnmarshalHosts - Unmarshal SSHProxy objects from a byte string.
func (k *Konnect) UnmarshalHosts(byteStr []byte) error {
	// Make a temporary type to hold host data.
	var tempHosts struct {
		Hosts map[string]interface{} `yaml:"hosts"`
	}

	// Unmarshal the byte string into the temp hosts struct.
	if err := yaml.Unmarshal(byteStr, &tempHosts); err != nil {
		return err
	}

	// Iterate through the unmarshalled hosts, and create SSHProxy
	// objects based on the type that was unmarshalled.
	for key, val := range tempHosts.Hosts {
		switch val.(type) {

		// If the value is a string, then it means that an
		// SSHProxy.Host value was supplied only. In this case,
		// we create an SSHProxy with this value as the `Host`.
		case string:
			// Construct an SSHProxy object.
			proxy := proxy.New("", val.(string), 0, "")
			// Fill in values from global config.
			proxy.PopulateFromProxy(k.Global)
			// Assign to Konnect.
			k.Hosts[key] = proxy

		// If the value is an interfact map, then it means
		// that an SSHProxy was possibly defined in full. In
		// this case, we marshal the map to a byte string, and
		// unmarhsal the byte string into an SSHProxy object.
		case map[interface{}]interface{}:
			// Turn the value into a byte string.
			byteStr, _ := yaml.Marshal(val)
			// Construct an SSHProxy object.
			newProxy := proxy.Default()
			// Unmarshal the byte string into an SSHProxy object.
			err := yaml.Unmarshal(byteStr, newProxy)
			if err != nil {
				return err
			}
			// Fill in values from global config.
			newProxy.PopulateFromProxy(k.Global)
			// Assign to Konnect.
			k.Hosts[key] = newProxy

		default:
			return errors.New("Unknown type for temp host")
		}
	}

	return nil
}

// UnmarshalTasks - Unmarshal SSHTask objects from a byte string.
func (k *Konnect) UnmarshalTasks(byteStr []byte) error {
	// Make a temporary type to hold the task data.
	var tempTasks struct {
		Tasks map[string]string `yaml:"tasks"`
	}

	// Unmarshal the byteStr into the temp tasks struct.
	if err := yaml.Unmarshal(byteStr, &tempTasks); err != nil {
		return err
	}

	// Iterate through the unmarshalled tasks,
	// and create SSHTask objects.
	for key, val := range tempTasks.Tasks {
		// Construct an SSHTask object.
		newTask := task.New(key, val)

		// Assign to Konnect.
		k.Tasks[key] = newTask
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

	// Unmarshal global proxy config from a byte string.
	if err := k.UnmarshalGlobal(byteStr); err != nil {
		return err
	}

	// Unmarshal SSHProxy objects from a byte string.
	if err := k.UnmarshalHosts(byteStr); err != nil {
		return err
	}

	// Unmarshal SSHTask objects from a byte string.
	if err := k.UnmarshalTasks(byteStr); err != nil {
		return err
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
		fmt.Println(proxy.PrintStatus())
	}
}
