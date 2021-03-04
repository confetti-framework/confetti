package commands

import (
	"github.com/confetti-framework/contract/inter"
)

type ExampleCommand struct {
	FirstFlag string `short:"f" flag:"first" description:"Configure your first flag" required:"true"`
}

func (t ExampleCommand) Name() string {
	return "example:command"
}

func (t ExampleCommand) Description() string {
	return "You can adjust this command to your wishes."
}

func (t ExampleCommand) Handle(c inter.Cli) inter.ExitCode {
	c.Info("Value in fist flag: %s", t.FirstFlag)
	return inter.Success
}
