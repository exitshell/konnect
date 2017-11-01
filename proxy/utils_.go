package proxy

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

// New - Create a new SSHProxy object with given values.
func New(user, host string, port int, key string) *SSHProxy {
	return &SSHProxy{
		User: user,
		Host: host,
		Port: port,
		Key:  key,
	}
}

// NewGlobal - Create a new SSHProxy object
// that will be used as a global config.
func NewGlobal() *SSHProxy {
	return &SSHProxy{
		User: "",
		Host: "",
		Port: getDefaultPort(),
		Key:  getDefaultKey(),
	}
}

// Default test.
func Default() *SSHProxy {
	return &SSHProxy{}
}
