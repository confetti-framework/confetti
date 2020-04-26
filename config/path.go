package config

import (
	"github.com/lanvard/contract/inter"
	"os"
	"path/filepath"
	"runtime"
)

const pathSeparator = string(os.PathSeparator)

type basePath string

func NewPath() inter.BasePath {
	_, filename, _, _ := runtime.Caller(0)
	return basePath(filepath.Dir(filepath.Dir(filename)))
}

// Get the path to the application "app" directory.
func (basePath basePath) AppPath() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the base path of the Lanvard installation.
func (basePath basePath) BasePath() string {
	return string(basePath)
}

// Get the path to the bootstrap directory.
func (basePath basePath) BootstrapPath() string {
	return string(basePath) + pathSeparator + "bootstrap"
}

// Get the path to the ration files.
func (basePath basePath) ConfigPath() string {
	return string(basePath) + pathSeparator + "app"
}

// Get the path to the database directory.
func (basePath basePath) DatabasePath() string {
	return string(basePath) + pathSeparator + "database"
}

// Get the path to the language files.
func (basePath basePath) LangPath() string {
	return basePath.ResourcePath() + pathSeparator + "lang"
}

// Get the path to the public / web directory.
func (basePath basePath) PublicPath() string {
	return string(basePath) + pathSeparator + "public"
}

// Get the path to the storage directory.
func (basePath basePath) StoragePath() string {
	return string(basePath) + pathSeparator + "storage"
}

// Get the path to the resources directory.
func (basePath basePath) ResourcePath() string {
	return string(basePath) + pathSeparator + "resources"
}

// Get the path to the environment file.
func (basePath basePath) EnvironmentFile() string {
	return string(basePath) + pathSeparator + ".env"
}

// Get the path to the environment file for environment testing.
func (basePath basePath) EnvironmentTestingFile() string {
	return string(basePath) + pathSeparator + ".env.testing"
}
