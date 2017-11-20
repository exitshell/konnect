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
	// Extra args.
	ExtraArgs string `yaml:"-"`
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
func (s *SSHProxy) PrintStatus() string {
	status := color.New(color.FgRed, color.Bold).SprintFunc()
	connectionStr := "FAIL"

	if s.Connection == true {
		status = color.New(color.FgCyan, color.Bold).SprintFunc()
		connectionStr = "OK"
	}
	return fmt.Sprintf("Connection %v\t-> [%v]", status(connectionStr), s.Name)
}

// Args - Return the full SSH command as a string slice of args.
func (s *SSHProxy) Args() []string {
	args := []string{
		"ssh",
		"-i",
		s.Key,
		"-p",
		strconv.Itoa(s.Port),
		fmt.Sprintf("%v@%v", s.User, s.Host),
	}
	// If SSHProxy has extra args, then
	// add them to the `args` slice.
	if len(s.ExtraArgs) > 0 {
		args = append(args, strings.Fields(s.ExtraArgs)...)
	}
	return args
}

// Connect to host.
func (s *SSHProxy) Connect() error {
	args := s.Args()
	fmt.Printf("\nWould run: %v\n", strings.Join(args, " "))

	// Make command, and pipe streams.
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run command.
	return cmd.Run()
}

// Validate SSHProxy fields.
func (s *SSHProxy) Validate() error {
	// Check that the `User` is not blank.
	if s.User == "" {
		return errors.New("User cannot be empty")
	}

	// Check that the `Host` is not blank.
	if s.Host == "" {
		return errors.New("Host cannot be empty")
	}

	// Expand key path if path contains tilde.
	expandedKey, err := tilde.Expand(s.Key)
	if err != nil {
		return fmt.Errorf("Cannot parse path %v", s.Key)
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
			return fmt.Errorf("Relative path error for %v, %v", dir, expandedKey)
		}
	}
	// Set `Key`.
	s.Key = expandedKey
	return nil
}

// PopulateFromProxy - Fill in values from a global proxy object.
func (s *SSHProxy) PopulateFromProxy(global *SSHProxy) {
	if s.User == "" {
		s.User = global.User
	}

	if s.Host == "" {
		s.Host = global.Host
	}

	if s.Port == 0 {
		s.Port = global.Port
	}

	if s.Key == "" {
		s.Key = global.Key
	}
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
