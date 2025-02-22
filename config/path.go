package config

import (
	"os"
	"path/filepath"
)

// Path contains a list of all paths you can use in your application.
var Path = struct {
	Base, Separator, ResultStorage, Storage string
}{
	/*
		|--------------------------------------------------------------------------
		| Base Path
		|--------------------------------------------------------------------------
		|
		| The base path is the fully qualified path to the root of your project.
		| Feel free to adjust this so that it fits to your needs.
		|
	*/
	Base:          basePath(),
	Separator:     string(os.PathSeparator),
	Storage:       filepath.Join(basePath(), "storage"),
	ResultStorage: filepath.Join(basePath(), "result"),
}

func basePath() string {
	root, _ := os.Getwd()
	return root
}
