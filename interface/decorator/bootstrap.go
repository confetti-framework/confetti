package decorator

import "laravelgo/interface/application"

type Bootstrap interface {
	Bootstrap(app application.App) application.App
}
