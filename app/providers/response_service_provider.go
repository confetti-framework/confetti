package providers

import (
	"github.com/confetti-framework/contract/inter"
	decorator "github.com/confetti-framework/foundation/decorator/response_decorator"
	"github.com/confetti-framework/foundation/encoder"
	net "net/http"
	"src/resources/views"
)

// ResponseServiceProvider is responsible for generate the correct response.
type ResponseServiceProvider struct{}

// Register is responsible for returning the correct response.
func (c ResponseServiceProvider) Register(container inter.Container) inter.Container {
	// Response decorators are responsible for modifying the response object.
	// All these decorators will be used.
	container.Bind("response_decorators", []inter.ResponseDecorator{
		decorator.LogError{},
		decorator.FilterSensitiveError{},
		decorator.HttpStatus{ErrorDefault: net.StatusInternalServerError},
		// add your custom decorators here
	})

	// Outcome encoders are responsible for converting an object
	// to a string. One encoder will be used.
	container.Bind("outcome_html_encoders", []inter.Encoder{
		// add your custom HTML encoders here
		encoder.ViewToHtml{},
		encoder.ErrorsToHtml{View: views.Error},
		encoder.StringerToHtml{},
		encoder.RawToHtml{},
		encoder.InterfaceToHtml{},
	})

	container.Bind("outcome_json_encoders", []inter.Encoder{
		// add your custom JSON encoders here
		encoder.JsonReaderToJson{},
		encoder.ErrorsToJson{},
		encoder.RawToJson{},
		encoder.JsonToJson{},
		encoder.InterfaceToJson{},
	})

	container.Bind("outcome_content_encoders", []inter.Encoder{
		// add your custom Content encoders here
		encoder.StringToString{},
	})

	return container
}
