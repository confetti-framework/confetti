package config

import (
	"os"
	"path/filepath"
)

// Path contains a list of all paths you can use in your application.
var Path = struct {
	//	Separator, Base, App, Service, Bootstrap, Config, Database, Public,

	//	Resource, SharedResource, SharedResourceCurrentService, Lang, Views string
	Base, ResultStorage, Storage string
}{
	//    Separator: string(os.PathSeparator),
	/*
		|--------------------------------------------------------------------------
		| Base Path
		|--------------------------------------------------------------------------
		|
		| The base path is the fully qualified path to the root of your project.
		| Feel free to adjust this so that it fits to your needs.
		|
	*/
	Base: basePath(),
	//	App:                          filepath.Join(basePath(), "app"),
	//	Service:                      env.String("APP_SERVICE"),
	//	Bootstrap:                    filepath.Join(basePath(), "bootstrap"),
	//	Config:                       filepath.Join(basePath(), "app"),
	//	Database:                     filepath.Join(basePath(), "database"),
	//	Public:                       filepath.Join(basePath(), "public"),
	Storage:       filepath.Join(basePath(), "storage"),
	ResultStorage: filepath.Join(basePath(), "result"),
	//	Resource:                     filepath.Join(basePath(), "resources"),
	//	SharedResource:               env.String("SHARED_RESOURCE_PATH"),
	//	SharedResourceCurrentService: filepath.Join(env.String("SHARED_RESOURCE_PATH"), env.String("APP_SERVICE")),
}

func basePath() string {
	root, _ := os.Getwd()
	return root
}
