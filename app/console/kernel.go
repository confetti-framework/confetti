package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"src/app/console/commands"
	"src/app/console/getters"
)

var flagGetters = func() []flag.Getter {
	return []flag.Getter{
		new(getters.StringList),
		new(getters.IntList),
	}
}

func NewKernel(app inter.App) console.Kernel {
	return console.Kernel{
		App: app,
		Commands: []inter.Command{
			commands.AppServe{},
			console.LogClear{},
			commands.ExampleCommand{},
		},
		FlagProviders: []func() []flag.Getter{flagGetters},
	}
}
