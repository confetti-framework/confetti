package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/foundation"
	"lanvard/http"
	pipeline "lanvard/src/app/http/middleware"
	"lanvard/src/bootstrap"
	net "net/http"
	"strings"
	"testing"
)

type PipeOneStruct struct {
	App foundation.Application
}

func (p PipeOneStruct) Handle(request pipeline.Passable, next pipeline.Destination) http.Response {
	content := request.Content()
	request = request.SetContent(content + " - before.one request ")

	response := next(request)

	return response.SetContent(response.Content() + " - after.one ")
}

type PipeTwoStruct struct {
	App foundation.Application
}

func (p PipeTwoStruct) Handle(request pipeline.Passable, next pipeline.Destination) http.Response {
	content := request.Content()
	request = request.SetContent(content + " - before.two request ")

	response := next(request)

	return response.SetContent(response.Content() + " - after.two ")
}

type PipeThreeStruct struct {
	App foundation.Application
}

func (p PipeThreeStruct) Handle(request pipeline.Passable, next pipeline.Destination) http.Response {
	content := request.Content()
	request = request.SetContent(content + " - before.three request ")

	response := next(request)

	return response.SetContent(response.Content() + " - after.three ")
}

func Test_normal_pipe(t *testing.T) {
	app := bootstrap.NewApp()
	requestOriginal, _ := net.NewRequest("GET", "/health-check", strings.NewReader("start"))
	request := http.NewRequest(app, *requestOriginal)

	middleware := pipeline.NewPipeline(app)

	var pipes []pipeline.PipeInterface
	pipes = append(pipes, PipeOneStruct{app})
	pipes = append(pipes, PipeTwoStruct{app})
	pipes = append(pipes, PipeThreeStruct{app})

	function := func(request http.Request) http.Response {
		return http.NewResponse().SetContent(request.Content())
	}

	result := middleware.Send(request).Through(pipes).Then(function)

	assert.Equal(
		t,
		"start - before.one request  - before.two request  - before.three request  - after.three  - after."+
			"two  - after.one ",
		result.Content(),
	)
}
