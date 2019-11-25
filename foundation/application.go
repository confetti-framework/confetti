package foundation

import (
	"path/filepath"
	"runtime"
)

type Application struct {
	// The service container
	Container Container

	// The base path for the Laravel installation.
	BasePath BasePath

	// Indicates if the application has been bootstrapped before.
	HasBeenBootstrapped bool
}

func (a Application) Make(abstract interface{}) interface{} {
	return a.Container.Make(abstract)
}

// Set the base path for the application.
func (a *Application) SetBasePath() {
	_, filename, _, _ := runtime.Caller(0)

	basePath := filepath.Dir(filepath.Dir(filename))
	a.BasePath = BasePath(basePath)
	a.bindPathsInContainer()
}

// Bind all of the application paths in the container.
func (a *Application) bindPathsInContainer() {
	a.Container.Instance("path", a.BasePath.Path())
	a.Container.Instance("path.base", a.BasePath.BasePath())
	a.Container.Instance("path.lang", a.BasePath.LangPath())
	a.Container.Instance("path.config", a.BasePath.ConfigPath())
	a.Container.Instance("path.public", a.BasePath.PublicPath())
	a.Container.Instance("path.storage", a.BasePath.StoragePath())
	a.Container.Instance("path.database", a.BasePath.DatabasePath())
	a.Container.Instance("path.resources", a.BasePath.ResourcePath())
	a.Container.Instance("path.bootstrap", a.BasePath.BootstrapPath())
}
