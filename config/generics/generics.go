// +build generic

package generics

import (
	"lanvard/config"
	"lanvard/http"
	"lanvard/pipeline"
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

/*
   |--------------------------------------------------------------------------
   | Generics
   |--------------------------------------------------------------------------
   |
   | Struct: Specify the struct to create a generic yourself
   | SaveTo: The go file need to be in an empty directory
   | Vars: Here you can override the default types
   |
*/
var Generics = []Generic{
	{
		Struct: pipeline.Pipeline{},
		SaveTo: config.App.BasePath + "/src/app/http/middleware/pipeline.go",
		Vars: Vars{
			"Passable": (*http.Request)(nil),
			"Result":   (*http.Response)(nil),
		},
	},
}
