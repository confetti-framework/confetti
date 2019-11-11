package test

import (
	"fmt"
	// "github.com/stretchr/testify/assert"
	"lanvard/foundation"
	pipelineContract "lanvard/interface/pipeline"
	"lanvard/pipeline"
	"lanvard/src/bootstrap"
	"strings"
	"testing"
)

type LowerPipeStruct struct {
	App foundation.Application
}

func (l LowerPipeStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	data = strings.ToLower(data.(string))
	return next(data)
}

type PrefixPipeStruct struct {
	App foundation.Application
}

func (p PrefixPipeStruct) Handle(data interface{}, next func(next interface{}) interface{}) interface{} {
	response := next(data)

	fmt.Println("response")
	fmt.Println(response)
	response = "test" + response.(string)
	return response
}

func Test_normal_pipe(t *testing.T) {
	middleware := pipeline.Pipeline(bootstrap.App())

	var pipes []pipelineContract.Pipe
	pipes = append(pipes, LowerPipeStruct{bootstrap.App()})
	pipes = append(pipes, PrefixPipeStruct{bootstrap.App()})

	function := func(data interface{}) interface{} {
		return data
	}

	result := middleware.Send("QWERTY").Through(pipes).Then(function)
	fmt.Println(result)

	// assert.Same(t, "qwerty", result)
}
