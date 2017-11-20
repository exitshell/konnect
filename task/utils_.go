package task

// New - Create a new SSHTask object with given values.
func New(name, command string) *SSHTask {
	return &SSHTask{
		Name:    name,
		Command: command,
	}
}
