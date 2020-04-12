package decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/decorator"
)

var bootstraps = []inter.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(app *foundation.Application) *foundation.Application {
	dec := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
