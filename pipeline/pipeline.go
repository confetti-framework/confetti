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

	var callbacks []func(data interface{}) interface{}
	var nextCallback = 0

	pipes := reverse(p.Pipes)

	for i, pipe := range pipes {
		pipe := pipe
		if i == 0 {
			callback := func(data interface{}) interface{} {
				return pipe.Handle(data, destination)
			}
			callbacks = append(callbacks, callback)
		} else {
			callback := func(data interface{}) interface{} {
				nextCallback--
				return pipe.Handle(data, callbacks[nextCallback])
			}
			callbacks = append(callbacks, callback)
		}
	}

	nextCallback = len(callbacks) - 1

	return callbacks[nextCallback](p.Passable)
}

func reverse(pipes []contract.Pipe) []contract.Pipe {
	for left, right := 0, len(pipes)-1; left < right; left, right = left+1, right-1 {
		pipes[left], pipes[right] = pipes[right], pipes[left]
	}

	return pipes
}
