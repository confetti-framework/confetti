package console

import (
	"flag"
	"github.com/confetti-framework/framework/foundation/console"
	"github.com/confetti-framework/framework/inter"
	"src/app/console/commands"
	"src/app/console/getters"
)

// NewKernel ensures that the kernel receives all existing commands
// and that the correct flag.Getters are used.
func NewKernel(app inter.App) console.Kernel {
	return console.Kernel{
		App: app,
		Commands: append(
			[]inter.Command{
				commands.ExampleCommand{},
				// Here you can add your own commands.
			}, console.FoundationCommands...,
		),

		// This list includes custom flag.Getters, you can create custom
		// types to cast flags from the command to a value.
		Getters: []flag.Getter{
			new(getters.StringList),
			new(getters.IntList),
		},
	}
}
