package bootstrap

import (
	"github.com/joho/godotenv"
	"lanvard/foundation"
)

type LoadEnvironmentVariables struct {
	environmentVariables map[string]string
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
