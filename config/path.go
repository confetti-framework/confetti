package config

import (
	"os"
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
	Separator, Base, App, Bootstrap, Config, Database, Public, Storage,
	Resource, Lang, Views string
}{
	Separator: separator,
	Base:      basePath(),
	App:       basePath() + separator + "app",
	Bootstrap: basePath() + separator + "bootstrap",
	Config:    basePath() + separator + "app",
	Database:  basePath() + separator + "database",
	Public:    basePath() + separator + "public",
	Storage:   basePath() + separator + "storage",
	Resource:  basePath() + separator + "resources",
	Lang:      basePath() + separator + "resources" + separator + "lang",
	Views:     basePath() + separator + "resources" + separator + "views",
}

func basePath() string {
	root, _ := os.Getwd()
	return root
}
