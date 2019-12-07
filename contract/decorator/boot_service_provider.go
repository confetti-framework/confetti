package decorator

import "lanvard/foundation"

type BootServiceProvider interface {
	Boot(app foundation.Application) foundation.Application
}
