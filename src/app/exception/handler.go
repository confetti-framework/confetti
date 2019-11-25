package exception

import "lanvard/foundation"

type Handler struct {
	app foundation.Application
}

func NewHandler(app foundation.Application) Handler {
	return Handler{app}
}
