package config

import (
	"github.com/confetti-framework/framework/foundation/loggers"
	"github.com/confetti-framework/framework/inter"
	"github.com/confetti-framework/framework/support/env"
	"github.com/confetti-framework/syslog/log_level"
	"os"
)

// Logging is responsible for all configuration related to logging
var Logging = struct {
	Default  string
	Channels map[string]inter.Logger
}{
	/*
	   |--------------------------------------------------------------------------
	   | Default Log Channel
	   |--------------------------------------------------------------------------
	   |
	   | This option defines the default log channel that gets used when writing
	   | messages to the logs. The name specified in this option should match
	   | one of the channels defined in the "Channels" configuration slice.
	   |
	*/
	Default: "stack",

	/*
	   |--------------------------------------------------------------------------
	   | Log Channels
	   |--------------------------------------------------------------------------
	   |
	   | Here you may configure the log channels for your application. Out of
	   | the box, Confetti uses the confetti-framework/syslog logging library. This gives
	   | you a variety of powerful log handlers / formatters to utilize.
	   |
	   | The name provided is for reference only, so you can log specifically to
	   | that channel. Feel free to compose your own channel.
	   |
	*/
	Channels: map[string]inter.Logger{
		"stack": loggers.Stack{
			Channels: []string{
				"daily",
				"stdout",
			},
		},

		"single": loggers.Syslog{
			Path:     Path.Storage + "/logs/default.log",
			MinLevel: log_level.DEBUG,
		},

		"daily": loggers.Syslog{
			Path:     Path.Storage + "/logs/{yyyy-mm-dd}_default.log",
			MinLevel: log_level.DEBUG,
			MaxFiles: 14,
		},

		"slack": loggers.Slack{
			WebhookUrl: env.StringOr("LOG_SLACK_WEBHOOK_URL", ""),
			MinLevel:   log_level.CRITICAL,
		},

		"stderr": loggers.Syslog{
			MinLevel: log_level.ERROR,
			Writer:   os.Stderr,
		},

		"stdout": loggers.Syslog{
			MinLevel: log_level.DEBUG,
			Writer:   os.Stdout,
		},
	},
}
