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
