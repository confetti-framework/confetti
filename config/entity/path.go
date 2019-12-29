package entity

import (
	"os"
	"path/filepath"
	"runtime"
)

const pathSeparator = string(os.PathSeparator)

type BasePath string

func NewBasePath() BasePath {
	_, filename, _, _ := runtime.Caller(0)
	return BasePath(filepath.Dir(filepath.Dir(filepath.Dir(filename))))
}

// Get the path to the application "app" directory.
func (basePath BasePath) AppPath() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the base path of the Lanvard installation.
func (basePath BasePath) BasePath() string {
	return string(basePath)
}

// Get the path to the bootstrap directory.
func (basePath BasePath) BootstrapPath() string {
	return string(basePath) + pathSeparator + "bootstrap"
}

// Get the path to the ration files.
func (basePath BasePath) ConfigPath() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the path to the database directory.
func (basePath BasePath) DatabasePath() string {
	return string(basePath) + pathSeparator + "database"
}

// Get the path to the language files.
func (basePath BasePath) LangPath() string {
	return basePath.ResourcePath() + pathSeparator + "lang"
}

// Get the path to the public / web directory.
func (basePath BasePath) PublicPath() string {
	return string(basePath) + pathSeparator + "public"
}

// Get the path to the storage directory.
func (basePath BasePath) StoragePath() string {
	return string(basePath) + pathSeparator + "storage"
}

// Get the path to the resources directory.
func (basePath BasePath) ResourcePath() string {
	return string(basePath) + pathSeparator + "resources"
}

// Get the path to the environment file.
func (basePath BasePath) EnvironmentFile() string {
	return string(basePath) + pathSeparator + ".env"
}

// Get the path to the environment file for environment testing.
func (basePath BasePath) EnvironmentTestingFile() string {
	return string(basePath) + pathSeparator + ".env.testing"
}
