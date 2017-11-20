package task

import (
	"fmt"
	"strings"
)

// SSHTask is a command that will be run on a specific host.
type SSHTask struct {
	Command string
	Name    string
}

// String representation of an SSHTask object.
func (t *SSHTask) String() string {
	return fmt.Sprintf("<SSHTask %v: %v>", t.Name, t.Command)
}

// Info - Return info for an SSHTask object.
func (t *SSHTask) Info() string {
	return fmt.Sprintf("[%v] %v\n", t.Name, t.Command)
}

// Args - Return the SSHTask command as a string slice.
func (t *SSHTask) Args() []string {
	return strings.Fields(t.Command)
}
