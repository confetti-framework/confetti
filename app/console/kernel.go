package console

import (
	"confetti-framework/app/console/commands"
	"confetti-framework/app/console/getters"
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
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
			commands.ExampleCommand{},
		},
		FlagProviders: []func() []flag.Getter{flagGetters},
	}
}
