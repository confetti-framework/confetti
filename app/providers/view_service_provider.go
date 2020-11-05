package providers

import (
	"github.com/lanvard/contract/inter"
	"html/template"
	"lanvard/config"
)

type ViewServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (v ViewServiceProvider) Register(container inter.Container) inter.Container {
	container.Singleton("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.ParseGlob(config.Path.Views + "/*")
	})

	return container
}
