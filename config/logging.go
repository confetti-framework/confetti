package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/support/env"
	"github.com/lanvard/syslog"
	"os"
)

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
	   | the box, Laravel uses the lanvard/syslog logging library. This gives
	   | you a variety of powerful log handlers / formatters to utilize.
	   |
	   | The name provided is for reference only, so you can log specifically to
	   | that channel. Feel free to compose your own channel.
	   |
	*/
	Channels: map[string]inter.Logger{
		"stack": loggers.Stack{
			Channels: []string{"daily"},
		},

		"single": loggers.Syslog{
			Path:     Path.Storage + "/logs/default.log",
			MinLevel: syslog.DEBUG,
		},

		"daily": loggers.Syslog{
			Path:     Path.Storage + "/logs/{yyyy-mm-dd}_default.log",
			MinLevel: syslog.DEBUG,
			MaxFiles: 14,
		},

		"slack": loggers.Slack{
			WebhookUrl: env.StringOr("LOG_SLACK_WEBHOOK_URL", ""),
			MinLevel:   syslog.CRIT,
		},

		"stderr": loggers.Syslog{
			MinLevel: syslog.ERR,
			Writer:   os.Stderr,
		},
	},
}
