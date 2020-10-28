package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var separator = string(os.PathSeparator)

/*
	|--------------------------------------------------------------------------
	| Base Path
	|--------------------------------------------------------------------------
	|
	| The base path is the fully qualified path to the root of your project.
	| Feel free to adjust this so that it fits to your needs.
	|
*/
var Path = struct {
	Separator, Base, App, Bootstrap, Config, Database, Public, Storage, Views,
	Resource, Lang, EnvironmentFile, EnvironmentTestingFile string
}{
	Separator:              separator,
	Base:                   basePath(),
	App:                    basePath() + separator + "app",
	Bootstrap:              basePath() + separator + "bootstrap",
	Config:                 basePath() + separator + "app",
	Database:               basePath() + separator + "database",
	Public:                 basePath() + separator + "public",
	Storage:                basePath() + separator + "storage",
	Views:                  basePath() + separator + "storage" + separator + "framework" + separator + "views",
	Resource:               basePath() + separator + "resources",
	Lang:                   basePath() + separator + "resources" + separator + "lang",
	EnvironmentFile:        basePath() + separator + ".env",
	EnvironmentTestingFile: basePath() + separator + ".env.testing",
}

func basePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}
