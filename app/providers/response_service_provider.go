package providers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"lanvard/config"
)

type ResponseServiceProvider struct{}

func (c ResponseServiceProvider) Register(container inter.Container) inter.Container {
	// Response decorators are responsible for modifying the response object.
	// All these decorators will be used.
	container.Bind("response_decorators", append(
		response_decorator.ResponseDecorators,
		// add your custom decorators here
	))

	// Outcome encoders are responsible for converting an object
	// to a string. One encoder will be used.
	container.Bind("outcome_html_encoders", append(
		outcome.HtmlEncoders,
		encoder.ErrorToHtml{TemplateFile: config.Path.Views + "/error.gohtml"},
		// add your custom HTML encoders here
	))
	container.Bind("outcome_json_encoders", append(
		outcome.JsonEncoders,
		// add your custom JSON encoders here
	))

	return container
}
