package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	// "github.com/stretchr/testify/assert"
	"lanvard/foundation"
	pipelineContract "lanvard/interface/pipeline"
	"lanvard/pipeline"
	"lanvard/src/bootstrap"
	"testing"
)

type PipeOneStruct struct {
	App foundation.Application
}

func (l PipeOneStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	data = data.(string) + " - bef.one data "
	println(" one data ")
	fmt.Println("before one:" + data.(string))
	response := next(data)
	fmt.Println("after one:" + response.(string))

	response = response.(string) + " - aft.one "
	return response
}

type PipeTwoStruct struct {
	App foundation.Application
}

func (p PipeTwoStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	data = data.(string) + " - bef.two data "
	println(" two data ")
	fmt.Println("before two:" + data.(string))
	response := next(data)
	fmt.Println("after two:" + response.(string))

	response = response.(string) + " - aft.two "
	return response
}

type PipeThreeStruct struct {
	App foundation.Application
}

func (p PipeThreeStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	data = data.(string) + " - bef.three data "
	println(" three data")
	fmt.Println("before three:" + data.(string))
	response := next(data)
	fmt.Println("after after:" + response.(string))

	response = response.(string) + " - aft.three "
	return response
}

func Test_normal_pipe(t *testing.T) {
	middleware := pipeline.Pipeline(bootstrap.App())

	var pipes []pipelineContract.Pipe
	pipes = append(pipes, PipeOneStruct{bootstrap.App()})
	pipes = append(pipes, PipeTwoStruct{bootstrap.App()})
	pipes = append(pipes, PipeThreeStruct{bootstrap.App()})

	function := func(data interface{}) interface{} {
		return data
	}

	result := middleware.Send("start").Through(pipes).Then(function)
	fmt.Println(result)

	assert.Equal(t, "qwerty", result.(string))
}
