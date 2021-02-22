package console

import (
	"confetti-framework/app/console/flag_type"
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
)

var flagGetters = func() []flag.Getter {
	return []flag.Getter{
		new(flag_type.StringList),
		new(flag_type.IntList),
	}
}

func NewKernel(app inter.App) console.Kernel {
	return console.Kernel{
		App: app,
		Commands: []inter.Command{
			AppServe{},
			YourFirstCommand{},
			console.LogClear{},
		},
		FlagProviders: []func() []flag.Getter{flagGetters},
	}
}
