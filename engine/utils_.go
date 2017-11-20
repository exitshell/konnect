package engine

import (
	"path/filepath"

	"github.com/exitshell/konnect/proxy"
	"github.com/exitshell/konnect/task"
)

// New - Create a new Konnect object.
func New() *Konnect {
	return &Konnect{
		Hosts:         make(map[string]*proxy.SSHProxy),
		Tasks:         make(map[string]*task.SSHTask),
		ProxyChan:     make(chan bool),
		CompletedChan: make(chan bool),
	}
}

// Init - Create a new Konnect object.
func Init(filename string) (*Konnect, error) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	konnect := New()
	if err = konnect.LoadFromFile(filename); err != nil {
		return nil, err
	}

	return konnect, nil
}
