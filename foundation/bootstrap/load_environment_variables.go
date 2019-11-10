package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"lanvard/foundation"
	"lanvard/interface/decorator"
)

type LoadEnvironmentVariablesStruct struct {
	environmentVariables map[string]string
}

func LoadEnvironmentVariables() decorator.Bootstrap {
	return LoadEnvironmentVariablesStruct{}
}

func (l LoadEnvironmentVariablesStruct) Bootstrap(app foundation.Application) foundation.Application {
	file := app.BasePath.EnvironmentFile()
	err := godotenv.Load(file)
	if err != nil {
		println(err)
		panic("Error loading " + file + " file in directory ")
	}

	return app
}
