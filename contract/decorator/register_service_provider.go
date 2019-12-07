package decorator

import "lanvard/foundation"

type RegisterServiceProvider interface {
	Register(app foundation.Application) foundation.Application
}
