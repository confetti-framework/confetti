package config

import (
	"github.com/lanvard/support/env"
	"os"
	"path/filepath"
	"runtime"
)

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
	Separator, Base, App, Bootstrap, Config, Database, Public, Storage,
	Resource, Lang, EnvironmentFile, EnvironmentTestingFile string
}{
	Separator:              pathSeparator(),
	Base:                   basePath(),
	App:                    basePath() + pathSeparator() + "app",
	Bootstrap:              basePath() + pathSeparator() + "bootstrap",
	Config:                 basePath() + pathSeparator() + "app",
	Database:               basePath() + pathSeparator() + "database",
	Public:                 basePath() + pathSeparator() + "public",
	Storage:                basePath() + pathSeparator() + "storage",
	Resource:               basePath() + pathSeparator() + "resources",
	Lang:                   basePath() + pathSeparator() + "resources" + pathSeparator() + "lang",
	EnvironmentFile:        basePath() + pathSeparator() + ".env",
	EnvironmentTestingFile: basePath() + pathSeparator() + ".env.testing",
}

func basePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}

func pathSeparator() string {
	return env.StringOr("PATH_SEPARATOR", string(os.PathSeparator))
}
