package console

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"io"
)

type YourFirstCommand struct {
	YourFirstFlag string `short:"f" flag:"first" description:"Configure your flag"`
}

func (t YourFirstCommand) Name() string {
	return "your:first-command"
}

func (t YourFirstCommand) Description() string {
	return "You can adjust this command to your wishes."
}

func (t YourFirstCommand) Handle(app inter.App, output io.Writer) inter.ExitCode {
	_, _ = fmt.Fprintln(output, "Value in first flag: "+t.YourFirstFlag)

	_, _ = fmt.Fprintln(output, "Done")
	return inter.Success
}
