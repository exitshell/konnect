package engine

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// Init - Create a new Konnect object.
func Init(filename string) *Konnect {
	filename, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}

	konnect := New()
	if err = konnect.LoadFromFile(filename); err != nil {
		log.Fatal(err)
	}

	return konnect
}

// List - Show info for all SSHProxy objects.
func (k *Konnect) List() {
	for _, proxy := range k.Hosts {
		fmt.Println(proxy.Info())
	}
}

// Args - Print SSH Args for a given host.
func (k *Konnect) Args(host string) {
	proxy, err := k.Get(host)
	if err != nil {
		log.Fatal(err)
	}

	argsStr := strings.Join(proxy.Args(), " ")
	fmt.Println(argsStr)
}

// Connect to host.
func (k *Konnect) Connect(host string) {
	proxy, err := k.Get(host)
	if err != nil {
		log.Fatal(err)
	}

	if err = proxy.Connect(); err != nil {
		log.Fatal(err)
	}
}

// Status - Check the status of one or more hosts.
func (k *Konnect) Status(hosts []string) {
	// Validate hosts. If a given host does not exist
	// in Konnect.Hosts, then throw an error.
	for _, host := range hosts {
		if _, ok := k.Hosts[host]; ok != true {
			log.Fatalf("[config] Undefined host %v\n", host)
		}
	}
	fmt.Printf("Getting status of hosts: %v\n", hosts)
}
