package exception

import (
	"github.com/lanvard/contract/inter"
)

type Handler struct {
	app inter.App
}

func NewHandler(app inter.App) Handler {
	return Handler{app}
}
