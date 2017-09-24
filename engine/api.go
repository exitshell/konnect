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
