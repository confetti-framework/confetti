package config

import (
	"lanvard/pipeline"
	"net/http"
)

type Vars map[string]interface{}

type ProfessorAble interface {
	Path() string
}

type Professor struct {
	Name   string
	Target ProfessorAble
	Vars   Vars
}

var Professors = []Professor{
	{
		Name:   "RequestPipeline",
		Target: pipeline.PipelineStruct{},
		Vars: Vars{
			"Passable": (*http.Request)(nil),
			"Result":   (*http.ResponseWriter)(nil),
		},
	},
}
