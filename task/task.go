package task

import "fmt"

// type SSHTask string
type SSHTask struct {
	Command string
	Name string
}

// String representation of an SSHTask object.
func (t *SSHTask) String() string {
	return fmt.Sprintf("<SSHTask %v: %v>", t.Name, t.Command)
}

// Info - Return info for an SSHTask object.
func (t *SSHTask) Info() string {
	return fmt.Sprintf("[%v]\n"+
		"  Command: %v\n",
	t.Name, t.Command)
}
