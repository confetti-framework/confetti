package config

import (
	"github.com/confetti-framework/routing"
	"confetti-framework/app/report"
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
		routing.MethodNotAllowedError,
		routing.RouteNotFoundError,
		report.ValidationError,
		report.NotFoundError,
	},
}
