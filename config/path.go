package config

import (
	"os"
	"path/filepath"
)

// Path contains a list of all paths you can use in your application.
var Path = struct {
	Separator, Base, App, Bootstrap, Config, Database, Public, Storage,
	Resource, Lang, Views string
}{
	Separator: string(os.PathSeparator),
	/*
		|--------------------------------------------------------------------------
		| Base Path
		|--------------------------------------------------------------------------
		|
		| The base path is the fully qualified path to the root of your project.
		| Feel free to adjust this so that it fits to your needs.
		|
	*/
	Base:      basePath(),
	App:       filepath.Join(basePath(), "app"),
	Bootstrap: filepath.Join(basePath(), "bootstrap"),
	Config:    filepath.Join(basePath(), "app"),
	Database:  filepath.Join(basePath(), "database"),
	Public:    filepath.Join(basePath(), "public"),
	Storage:   filepath.Join(basePath(), "storage"),
	Resource:  filepath.Join(basePath(), "resources"),
}

func basePath() string {
	root, _ := os.Getwd()
	return root
}
