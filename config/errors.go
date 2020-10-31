package config

import "lanvard/app/report"

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
		report.UserError,
	},
}
