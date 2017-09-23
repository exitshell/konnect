package sshproxy

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"

	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

// SSHProxy is a type which contains configs for an ssh connection.
type SSHProxy struct {
	// SSH configs.
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
	// Config filepath.
	Filename string
	// Name of SSHProxy config.
	Name string
}

// String representation of SSHProxy.
func (s *SSHProxy) String() string {
	if s.User == "" || s.Host == "" {
		return "<SSHProxy: empty>"
	}
	return fmt.Sprintf("<SSHProxy: %v@%v>", s.User, s.Host)
}

// Info for SSHProxy value.
func (s *SSHProxy) Info() string {
	return fmt.Sprintf("[%v]\n"+
		"  User: %v\n"+
		"  Host: %v\n"+
		"  Port: %v\n"+
		"  Key: %v\n",
		s.Name, s.User, s.Host, s.Port, s.Key)
}

// Validate SSHProxy fields.
func (s *SSHProxy) Validate() error {
	// Check that the `User` is not blank.
	if s.User == "" {
		return errors.New("[config] Invalid User")
	}

	// Check that the `Host` is not blank.
	if s.Host == "" {
		return errors.New("[config] Invalid Host")
	}

	// Expand key path if path contains tilde.
	expandedKey, err := tilde.Expand(s.Key)
	if err != nil {
		return errors.New("[config] Expanding Key path")
	}

	// If `Key` is not an absolute path, then make `Key` an
	// absolute path that is relative to `Filename`.
	if !path.IsAbs(expandedKey) {
		// Get dirname of config file.
		dir := filepath.Dir(s.Filename)
		// Join dirname and keypath.
		expandedKey = filepath.Join(dir, expandedKey)
		// Absolute path.
		expandedKey, err = filepath.Abs(expandedKey)
		if err != nil {
			errMsg := fmt.Sprintf("[config] Key path %v", err)
			return errors.New(errMsg)
		}
	}
	// Set `Key`.
	s.Key = expandedKey
	return nil
}

// UnmarshalYAML - Populate an SSHProxy struct from a yaml byte string.
// https://goo.gl/yvLJkj
func (s *SSHProxy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Create an alias type.
	type SSHAlias SSHProxy

	// Convert new value to the alias type.
	var temp = (*SSHAlias)(Default())

	// Unmarshal into the aliased type.
	if err := unmarshal(temp); err != nil {
		return err
	}

	// Case the alias back to the original type.
	*s = SSHProxy(*temp)
	return nil
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

// Default - Create a new SSHProxy object with default values.
func Default() *SSHProxy {
	return &SSHProxy{
		User: "",
		Host: "",
		Port: getDefaultPort(),
		Key:  getDefaultKey(),
	}
}
