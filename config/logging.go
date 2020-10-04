package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/drivers"
	"github.com/lanvard/support/env"
	"github.com/sirupsen/logrus"
)

var Logging = struct {
	Default  string
	Channels map[string]inter.LogChannel
}{
	/*
	   |--------------------------------------------------------------------------
	   | Default Log Channel
	   |--------------------------------------------------------------------------
	   |
	   | This option defines the default log channel that gets used when writing
	   | messages to the logs. The name specified in this option should match
	   | one of the channels defined in the "channels" configuration array.
	   |
	*/
	Default: "stack",

	/*
	   |--------------------------------------------------------------------------
	   | Log Channels
	   |--------------------------------------------------------------------------
	   |
	   | Here you may configure the log channels for your application. Out of
	   | the box, Laravel uses the Monolog PHP logging library. This gives
	   | you a variety of powerful log handlers / formatters to utilize.
	   |
	   | Available Drivers: "single", "daily", "slack", "syslog",
	   |                    "errorlog",
	   |                    "custom", "stack"
	   |
	*/
	Channels: map[string]inter.LogChannel{
		"stack": drivers.Stack{
			Channels: []string{"single"},
		},

		"single": drivers.LogRus{
			Path:  Path.Storage + "/logs/lanvard.log",
			Level: logrus.DebugLevel,
		},

		"daily": drivers.LogRus{
			Path:  Path.Storage + "/logs/lanvard.log",
			Level: logrus.DebugLevel,
			Days:  14,
		},

		"slack": drivers.Slack{
			Url:      env.StringOr("LOG_SLACK_WEBHOOK_URL", ""),
			Username: "Lanvard log",
			Emoji:    ":boom:",
			Level:    logrus.ErrorLevel,
		},
	},
}
