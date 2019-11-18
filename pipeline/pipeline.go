// This file is a template to generate generics
// This is the original file
// +build generic

package pipeline

import (
	"lanvard/foundation"
	"lanvard/support/caller"
)

type Passable = interface{}
type Result = interface{}

type Destination func(data Passable) Result
type PipeInterface interface {
	Handle(data Passable, next Destination) Result
}

// noinspection GoNameStartsWithPackageName
type PipelineStruct struct {
	App foundation.Application

	// The object being passed through the contract.
	Passable Passable

	// the array of pipes.
	Pipes []PipeInterface
}

func Pipeline(app foundation.Application) PipelineStruct {
	return PipelineStruct{App: app}
}

func (p PipelineStruct) Path() string {
	return caller.Path()
}

// Set the object being sent through the contract.
func (p PipelineStruct) Send(passable Passable) PipelineStruct {
	p.Passable = passable

	return p
}

// Set the array of pipes.
func (p PipelineStruct) Through(pipes []PipeInterface) PipelineStruct {
	p.Pipes = pipes

	return p
}

// Run the contract with a final destination callback.
func (p PipelineStruct) Then(destination Destination) Result {

	var callbacks []func(data Passable) Result
	var nextCallback = 0

	pipes := reverse(p.Pipes)

	for i, pipe := range pipes {
		pipe := pipe
		if i == 0 {
			callback := func(data Passable) Result {
				return pipe.Handle(data, destination)
			}
			callbacks = append(callbacks, callback)
		} else {
			callback := func(data Passable) Result {
				nextCallback--
				return pipe.Handle(data, callbacks[nextCallback])
			}
			callbacks = append(callbacks, callback)
		}
	}

	nextCallback = len(callbacks) - 1

	return callbacks[nextCallback](p.Passable)
}

func reverse(pipes []PipeInterface) []PipeInterface {
	for left, right := 0, len(pipes)-1; left < right; left, right = left+1, right-1 {
		pipes[left], pipes[right] = pipes[right], pipes[left]
	}

	return pipes
}
