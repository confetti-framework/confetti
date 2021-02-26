package commands

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"io"
)

type ExampleCommand struct {
	FirstFlag string `short:"f" flag:"first" description:"Configure your first flag" required:"true"`
	SecondFlag bool `short:"s" flag:"second" description:"Configure your second flag" required:"true"`
	Ids []int `short:"I"`
}

func (t ExampleCommand) Name() string {
	return "example:command"
}

func (t ExampleCommand) Description() string {
	return "You can adjust this command to your wishes."
}

func (t ExampleCommand) Handle(app inter.App, output io.Writer) inter.ExitCode {
	_, _ = fmt.Fprintln(output, "Value in fist flag: "+t.FirstFlag)
	support.Dump(t.Ids)

	if t.SecondFlag {
		_, _ = fmt.Fprintln(output, "Value in second flag is true")
	}

	_, _ = fmt.Fprintln(output, "Done")
	return inter.Success
}
