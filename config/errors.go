package config

import (
	"github.com/lanvard/routing"
	"lanvard/app/report"
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
