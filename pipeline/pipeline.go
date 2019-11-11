package pipeline

import (
	"lanvard/foundation"
	contract "lanvard/interface/pipeline"
)

// noinspection GoNameStartsWithPackageName
type PipelineStruct struct {
	App foundation.Application

	// The object being passed through the contract.
	Passable interface{}

	// the array of pipes.
	Pipes []contract.Pipe
}

func Pipeline(app foundation.Application) PipelineStruct {
	return PipelineStruct{App: app}
}

// Set the object being sent through the contract.
func (p PipelineStruct) Send(passable interface{}) PipelineStruct {
	p.Passable = passable

	return p
}

// Set the array of pipes.
func (p PipelineStruct) Through(pipes []contract.Pipe) PipelineStruct {
	p.Pipes = pipes

	return p
}

// Run the contract with a final destination callback.
func (p PipelineStruct) Then(destination func(data interface{}) interface{}) interface{} {

	var callback func(data interface{}) interface{}
	// var lastCallback func(data interface{}) interface{}
	var pipe contract.Pipe

	// var destinationPipe = Destination(p.App, p.Passable, destination)

	var destinationCallback = func(data interface{}) interface{} {
		return destination(data)
	}
	callback = destinationCallback

	for _, pipe = range p.Pipes {

		callback = func(data interface{}) interface{} {
			// next := p.Pipes[len(p.Pipes)-1-i]

			return pipe.Handle(p.Passable, destinationCallback)
		}
		// lastCallback = callback

		// fmt.Println(i)
		// fmt.Println(pipe)
		// next := p.Pipes[len(p.Pipes)-1-i]

		// handle = p.Pipes[len(p.Pipes)-1-i].Handle(p.Passable)
	}

	return callback(p.Passable)
}

// func (p PipelineStruct) response() http.ResponseWriter {
// 	return p.app.Make((*http.ResponseWriter)(nil)).(http.ResponseWriter)
// }
