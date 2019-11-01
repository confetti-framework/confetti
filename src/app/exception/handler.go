package exception

import "lanvard/foundation"

type HandlerStruct struct {
	app foundation.Application
}

func Handler(app foundation.Application) HandlerStruct {
	return HandlerStruct{app}
}
