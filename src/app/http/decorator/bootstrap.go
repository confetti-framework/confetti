package decorator

import (
	contract "github.com/lanvard/contract/decorator"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/bootstrap"
	"github.com/lanvard/foundation/decorator"
)

var bootstraps = []contract.Bootstrap{
	bootstrap.BasePathProvider{},
	bootstrap.LoadEnvironmentVariables{},
	RegisterProviders{},
	BootProviders{},
}

func Bootstrap(app foundation.Application) foundation.Application {
	dec := decorator.BootstrapDecorator{Bootstraps: bootstraps}

	return dec.BootstrapWith(app)
}
