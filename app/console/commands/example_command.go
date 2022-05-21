package commands

import (
	"github.com/confetti-framework/framework/inter"
)

// ExampleCommand to give you an example of what a command might look like.
type ExampleCommand struct {
	FirstFlag string `short:"f" flag:"first" description:"Configure your first flag." required:"true"`
}

// Name of the command
func (t ExampleCommand) Name() string {
	return "example:command"
}

// Description of the command
func (t ExampleCommand) Description() string {
	return "You can adjust this command to your wishes."
}

// Handle contains the logic of the command
func (t ExampleCommand) Handle(c inter.Cli) inter.ExitCode {
	c.Info("Value in fist flag: %s", t.FirstFlag)
	return inter.Success
}
