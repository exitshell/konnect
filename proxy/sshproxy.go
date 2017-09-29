package proxy

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	color "github.com/fatih/color"
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
	Filename string `yaml:"-"`
	// Name of SSHProxy config.
	Name string `yaml:"-"`
	// A bool to determine if the connection is ok.
	Connection bool `yaml:"-"`
}

// String representation of an SSHProxy object.
func (s *SSHProxy) String() string {
	if s.User == "" || s.Host == "" {
		return "<SSHProxy: empty>"
	}
	return fmt.Sprintf("<SSHProxy: %v@%v>", s.User, s.Host)
}

// Info - Return info for an SSHProxy object.
func (s *SSHProxy) Info() string {
	return fmt.Sprintf("[%v]\n"+
		"  User: %v\n"+
		"  Host: %v\n"+
		"  Port: %v\n"+
		"  Key: %v\n",
		s.Name, s.User, s.Host, s.Port, s.Key)
}

// PrintStatus - Return connection status for an SSHProxy object.
func (s *SSHProxy) PrintStatus() {
	status := color.New(color.FgRed, color.Bold).SprintFunc()
	connectionStr := "FAIL"

	if s.Connection == true {
		status = color.New(color.FgCyan, color.Bold).SprintFunc()
		connectionStr = "OK"
	}
	fmt.Printf("Connection %v\t-> [%v]\n", status(connectionStr), s.Name)
}

// Args - Return the full SSH command for SSHProxy.
func (s *SSHProxy) Args() []string {
	return []string{
		"ssh",
		"-i",
		s.Key,
		"-p",
		strconv.Itoa(s.Port),
		fmt.Sprintf("%v@%v", s.User, s.Host),
	}
}

// Connect to host.
func (s *SSHProxy) Connect() error {
	args := s.Args()

	// Make command, and pipe streams.
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run command.
	return cmd.Run()
	// if err := cmd.Run(); err != nil {
	// 	return err
	// }
	// return nil
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

// TestConnection - Test SSHProxy connection.
func (s *SSHProxy) TestConnection() {
	// Create timeout command context.
	// https://goo.gl/8dPsth
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	sshArgs := s.Args()
	sshArgs = append(sshArgs, "echo ping")

	// Create command with context.
	cmd := exec.CommandContext(ctx, sshArgs[0], sshArgs[1:]...)
	out, err := cmd.Output()

	// Check if the command timed-out.
	if ctx.Err() == context.DeadlineExceeded {
		s.Connection = false
		return
	}

	// Handle if error occured.
	if err != nil {
		s.Connection = false
		return
	}

	// Convert byte slice to string.
	outStr := string(out)
	// Trim newline chars from output str.
	// https://stackoverflow.com/a/44449581
	outStr = strings.TrimRight(outStr, "\r\n")

	// Handle if the result is not what is expected.
	if outStr != "ping" {
		s.Connection = false
		return
	}

	s.Connection = true
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
