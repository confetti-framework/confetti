package providers

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type PathServiceProvider struct{}

// Register any container services.
func (a PathServiceProvider) Register(container inter.Container) inter.Container {

	path := config.App.BasePath
	container.Instance("path.app", path.AppPath())
	container.Instance("path.base", path.BasePath())
	container.Instance("path.lang", path.LangPath())
	container.Instance("path.config", path.ConfigPath())
	container.Instance("path.public", path.PublicPath())
	container.Instance("path.storage", path.StoragePath())
	container.Instance("path.database", path.DatabasePath())
	container.Instance("path.resources", path.ResourcePath())
	container.Instance("path.bootstrap", path.BootstrapPath())

	return container
}
