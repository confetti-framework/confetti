package pipeline

import (
	"fmt"
	"lanvard/foundation"
)

type DestinationStruct struct {
	App         foundation.Application
	passable    interface{}
	destination func(data interface{}) interface{}
}

func Destination(app foundation.Application, passable interface{}, destination func(data interface{}) interface{}) DestinationStruct {
	return DestinationStruct{app, passable, destination}
}

func (d DestinationStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	fmt.Println("yoooo")
	return d.destination(data)
}
