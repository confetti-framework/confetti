package generics

import (
	"lanvard/config"
	"lanvard/pipeline"
	"net/http"
)

type Vars map[string]interface{}

type GenericAble interface {
	Path() string
}

type Generic struct {
	Struct GenericAble
	SaveTo string
	Vars   Vars
}

var Generics = []Generic{
	{
		// Specify the struct to create a generic yourself
		Struct: pipeline.PipelineStruct{},

		// The go file need to be in an empty directory
		SaveTo: config.App.BasePath + "/src/app/http/pipeline.go",

		// Here you can override the default types
		Vars: Vars{
			"Passable": (*http.Request)(nil),
			"Result":   (*http.ResponseWriter)(nil),
		},
	},
}
