package providers

import (
	"github.com/lanvard/contract/inter"
	"html/template"
	"lanvard/config"
)

type ViewServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (v ViewServiceProvider) Register(container inter.Container) inter.Container {
	container.Singleton("template_builder", func(templateBuilder *template.Template) (*template.Template, error) {
		if t, _ := templateBuilder.ParseGlob(config.Path.Views + "/*/*/*/*/*.gohtml"); t != nil {
			templateBuilder = t
		}
		if t, _ := templateBuilder.ParseGlob(config.Path.Views + "/*/*/*/*.gohtml"); t != nil {
			templateBuilder = t
		}
		if t, _ := templateBuilder.ParseGlob(config.Path.Views + "/*/*/*.gohtml"); t != nil {
			templateBuilder = t
		}
		if t, _ := templateBuilder.ParseGlob(config.Path.Views + "/*/*.gohtml"); t != nil {
			templateBuilder = t
		}
		return templateBuilder.ParseGlob(config.Path.Views + "/*.gohtml")
	})

	return container
}
