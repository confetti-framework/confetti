package exception

import "laravelgo/foundation"

type HandlerStruct struct {
	app foundation.ApplicationStruct
}

func Handler(app foundation.ApplicationStruct) HandlerStruct {
	return HandlerStruct{app}
}
