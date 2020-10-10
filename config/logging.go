package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
)

var Logging = struct {
	Default string
	Loggers map[string]inter.Logger
}{
	/*
	   |--------------------------------------------------------------------------
	   | Default Logger
	   |--------------------------------------------------------------------------
	   |
	   | This option defines the default logger that gets used when writing
	   | messages to the logs. The name specified in this option should match
	   | one of the loggers defined in the "Loggers" configuration.
	   |
	*/
	Default: "stack",

	/*
	   |--------------------------------------------------------------------------
	   | Log Loggers
	   |--------------------------------------------------------------------------
	   |
	   | Here you may configure the log loggers for your application. Out of
	   | the box, Laravel uses the lanvard/logrus logging library. This gives
	   | you a variety of powerful log handlers / formatters to utilize.
	   |
	   | The given key is for reference only. Feel free to compose your own logger.
	   |
	*/
	Loggers: map[string]inter.Logger{
		"stack": loggers.Stack{
			Loggers: []string{"daily"},
		},

		"single": loggers.Syslog{
			Path:     Path.Storage + "/logs/lanvard.log",
			MinLevel: syslog.INFO,
			AppName:  App.Name,
		},

		"daily": loggers.Syslog{
			Path:     Path.Storage + "/logs/{yyyy-mm-dd}_lanvard.log",
			MinLevel: syslog.INFO,
			AppName:  App.Name,
			MaxFiles: 14,
		},

		// "slack": loggers.Slack{
		// 	Url:      env.StringOr("LOG_SLACK_WEBHOOK_URL", ""),
		// 	Username: "Lanvard log",
		// 	Emoji:    ":boom:",
		// 	Level:    logrus.ErrorLevel,
		// },
	},
}
