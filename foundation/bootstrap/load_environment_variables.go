package bootstrap

import (
	"github.com/joho/godotenv"
	"lanvard/foundation"
	"lanvard/interface/decorator"
)

type LoadEnvironmentVariables struct {
	environmentVariables map[string]string
}

func BootLoadEnvironmentVariables() decorator.Bootstrap {
	return LoadEnvironmentVariables{}
}

func (l LoadEnvironmentVariables) Bootstrap(app foundation.Application) foundation.Application {
	file := app.BasePath.EnvironmentFile()
	err := godotenv.Load(file)
	if err != nil {
		println(err)
		panic("Error loading " + file + " file in directory ")
	}

	return app
}
