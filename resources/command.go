package resources

import (
	"io"
	"os/exec"
)

type Command struct {
	Command string
	Output  io.Writer
	Args    []string
}

func NewCommand(command string, args ...string) *Command {
	s := Command{
		Command: command,
		Args:    []string{},
	}

	return &s
}

func (s *Command) SetArgs(args ...string) *Command {
	s.Args = append(s.Args, args...)

	return s
}

func (s *Command) SetOutput(w io.Writer) *Command {
	s.Output = w

	return s
}

func (s *Command) Apply() error {
	cmd := exec.Command(s.Command, s.Args...)
	if s.Output != nil {
		cmd.Stdout = s.Output
	}

	return cmd.Run()
}

func (p *Command) Remove() error { return nil }
