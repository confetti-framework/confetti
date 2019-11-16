package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/foundation"
	"lanvard/pipeline"
	"lanvard/src/bootstrap"
	"testing"
)

type PipeOneStruct struct {
	App foundation.Application
}

func (p PipeOneStruct) Handle(data pipeline.Passable, next pipeline.Destination) pipeline.Result {
	data = data.(string) + " - before.one data "
	response := next(data)

	response = response.(string) + " - after.one "
	return response
}

type PipeTwoStruct struct {
	App foundation.Application
}

func (p PipeTwoStruct) Handle(data pipeline.Passable, next pipeline.Destination) pipeline.Result {
	data = data.(string) + " - before.two data "
	response := next(data)

	response = response.(string) + " - after.two "
	return response
}

type PipeThreeStruct struct {
	App foundation.Application
}

func (p PipeThreeStruct) Handle(data pipeline.Passable, next pipeline.Destination) pipeline.Result {
	data = data.(string) + " - before.three data "
	response := next(data)

	response = response.(string) + " - after.three "
	return response
}

func Test_normal_pipe(t *testing.T) {
	app := bootstrap.App()
	middleware := pipeline.Pipeline(app)

	var pipes []pipeline.PipeInterface
	pipes = append(pipes, PipeOneStruct{app})
	pipes = append(pipes, PipeTwoStruct{app})
	pipes = append(pipes, PipeThreeStruct{app})

	function := func(data interface{}) interface{} {
		return data
	}

	result := middleware.Send("start").Through(pipes).Then(function)

	assert.Equal(t, "start - before.one data  - before.two data  - before.three data  - after.three  - after.two  - after.one ",
		result.(string))
}
