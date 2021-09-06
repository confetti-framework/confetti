package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"src/app/console/commands"
	"src/app/console/getters"
)

// flagGetters contains a list with custom flag.Getters, you can create custom
// types to cast the flags to a value.
var flagGetters = func() []flag.Getter {
	return []flag.Getter{
		new(getters.StringList),
		new(getters.IntList),
	}
}

// NewKernel ensures that the kernel receives all existing commands and that the
// correct flag.Getters are used.
func NewKernel(app inter.App) console.Kernel {
	return console.Kernel{
		App: app,
		Commands: []inter.Command{
			console.AppServe{},
			console.LogClear{},
			commands.ExampleCommand{},
			console.AppInfo{},
		},
		FlagProviders: []func() []flag.Getter{flagGetters},
	}
}
