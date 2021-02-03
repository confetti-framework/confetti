package config

import (
	"confetti-framework/app/report"
	routing2 "github.com/confetti-framework/foundation/http/routing"
)

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
		routing2.MethodNotAllowedError,
		routing2.RouteNotFoundError,
		report.ValidationError,
		report.NotFoundError,
	},
}
