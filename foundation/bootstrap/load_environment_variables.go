package bootstrap

import (
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

	err := godotenv.Load(app.BasePath.EnvironmentFile())
	if err != nil {
		println(err)
		panic("Error loading .env file")
	}

	return app
}
