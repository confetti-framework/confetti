package decorator

import (
	contract "github.com/lanvard/contract/decorator"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/decorator"
)

var bootstraps = []contract.Bootstrap{
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(app *foundation.Application) *foundation.Application {
	dec := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
