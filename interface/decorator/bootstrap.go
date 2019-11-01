package decorator

import (
	"lanvard/foundation"
)

type Bootstrap interface {
	Bootstrap(app foundation.Application) foundation.Application
}
