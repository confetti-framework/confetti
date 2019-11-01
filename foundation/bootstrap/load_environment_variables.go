package bootstrap

import (
	"github.com/joho/godotenv"
	"laravelgo/interface/application"
	"laravelgo/interface/decorator"
	"log"
)

type LoadEnvironmentVariablesStruct struct {
	environmentVariables map[string]string
}

func LoadEnvironmentVariables() decorator.Bootstrap {
	return LoadEnvironmentVariablesStruct{}
}

func (l LoadEnvironmentVariablesStruct) Bootstrap(app application.App) application.App {

	err := godotenv.Load()
	if err != nil {
		println(err)
		log.Fatal("Error loading .env file")
	}

	l.environmentVariables = environmentVariables

	println(environmentVariables)
	return app
}