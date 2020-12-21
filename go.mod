module confetti-framework

go 1.15

require (
	github.com/confetti-framework/contract v0.0.0
	github.com/confetti-framework/errors v0.10.0
	github.com/confetti-framework/foundation v0.2.2
	github.com/confetti-framework/routing v0.2.0
	github.com/confetti-framework/support v0.1.0
	github.com/confetti-framework/syslog v0.0.0
	github.com/confetti-framework/validation v0.0.0-20201116205703-8aad92bf4d9a
	golang.org/x/text v0.3.2
)

replace (
	github.com/confetti-framework/contract v0.0.0 => ../github.com/confetti-framework/contract
	github.com/confetti-framework/errors v0.10.0 => ../github.com/confetti-framework/errors
	github.com/confetti-framework/foundation v0.2.2 => ../github.com/confetti-framework/foundation
	github.com/confetti-framework/routing v0.2.0 => ../github.com/confetti-framework/routing
	github.com/confetti-framework/support v0.1.0 => ../github.com/confetti-framework/support
	github.com/confetti-framework/syslog v0.0.0 => ../github.com/confetti-framework/syslog
	github.com/confetti-framework/validation v0.0.0-20201116205703-8aad92bf4d9a => ../github.com/confetti-framework/validation
)
