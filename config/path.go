package config

import (
	"os"
	. "path/filepath"
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
	Resource, Lang, Views string
}{
	Separator: string(os.PathSeparator),
	Base:      basePath(),
	App:       Join(basePath(), "app"),
	Bootstrap: Join(basePath(), "bootstrap"),
	Config:    Join(basePath(), "app"),
	Database:  Join(basePath(), "database"),
	Public:    Join(basePath(), "public"),
	Storage:   Join(basePath(), "storage"),
	Resource:  Join(basePath(), "resources"),
	Lang:      Join(basePath(), "resources", "lang"),
	Views:     Join(basePath(), "resources", "views"),
}

func basePath() string {
	root, _ := os.Getwd()
	return root
}
