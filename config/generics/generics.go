// +build generic

package generics

import (
	"github.com/lanvard/foundation/http/lanvard"
	"github.com/lanvard/pipeline"
	"lanvard/config"
)

type Vars map[string]interface{}

type GenericAble interface {
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
		SaveTo: string(config.App.Path()) + "/app/http/middleware/pipeline.go",
		Vars: Vars{
			"Passable": (*lanvard.Request)(nil),
			"Result":   (*lanvard.Response)(nil),
		},
	},
}
