package providers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/validation/val_errors"
	"html/template"
	"src/config"
	"strings"
)

// ViewServiceProvider Register here your view helpers and templates to the `template.Template` builder.
type ViewServiceProvider struct{}

// Register your view helpers and templates to the `template.Template` builder.
func (v ViewServiceProvider) Register(container inter.Container) inter.Container {
	container.Bind("template_builder", func(templateBuilder *template.Template) (*template.Template, error) {
		templateBuilder = addFunctions(templateBuilder)
		return addTemplates(templateBuilder)
	})

	return container
}

// Register your view helpers.
func addFunctions(templateBuilder *template.Template) *template.Template {
	return templateBuilder.Funcs(template.FuncMap{
		"Replace": func(input, from, to string) string {
			return strings.Replace(input, from, to, -1)
		},
		"Trim":  strings.Trim,
		"Error": val_errors.FindError,
	})
}

// Register your view templates.
func addTemplates(templateBuilder *template.Template) (*template.Template, error) {
	if t, _ := templateBuilder.ParseFS(config.Embed.Views, "*/*/*/*/*.gohtml"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseFS(config.Embed.Views, "*/*/*/*.gohtml"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseFS(config.Embed.Views, "*/*/*.gohtml"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseFS(config.Embed.Views, "*/*.gohtml"); t != nil {
		templateBuilder = t
	}
	return templateBuilder.ParseFS(config.Embed.Views, "*.gohtml")
}
