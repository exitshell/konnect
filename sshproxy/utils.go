package sshproxy

import (
	"os/user"
	"path/filepath"
)

// Return the home directory.
func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// Return the absolute path of the default SSH key.
func getDefaultKey() string {
	defaultKey := ".ssh/id_rsa"
	return filepath.Join(getHomeDir(), defaultKey)
}

// Return the default port
func getDefaultPort() int {
	return 22
}
