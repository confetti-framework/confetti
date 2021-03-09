package config

import (
	"github.com/confetti-framework/foundation/http/routing"
	"src/app/report"
)

// Errors contains the configuration to handle the errors correctly
var Errors = struct {
	NoLogging []error
}{
	/*
	   |--------------------------------------------------------------------------
	   | No Logging
	   |--------------------------------------------------------------------------
	   |
	   | A list of the error types that should not be logged.
	   |
	*/
	NoLogging: []error{
		routing.MethodNotAllowedError,
		routing.RouteNotFoundError,
		report.ValidationError,
		report.NotFoundError,
	},
}
