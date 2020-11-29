module lanvard

go 1.15

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/errors v0.10.0
	github.com/lanvard/foundation v0.2.2
	github.com/lanvard/routing v0.2.0
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0
	github.com/lanvard/validation v0.0.0-20201116205703-8aad92bf4d9a
	golang.org/x/text v0.3.2
)

replace (
	github.com/lanvard/contract v0.0.0 => ../github.com/lanvard/contract
	github.com/lanvard/errors v0.10.0 => ../github.com/lanvard/errors
	github.com/lanvard/foundation v0.2.2 => ../github.com/lanvard/foundation
	github.com/lanvard/routing v0.2.0 => ../github.com/lanvard/routing
	github.com/lanvard/support v0.1.0 => ../github.com/lanvard/support
	github.com/lanvard/syslog v0.0.0 => ../github.com/lanvard/syslog
	github.com/lanvard/validation v0.0.0-20201116205703-8aad92bf4d9a => ../github.com/lanvard/validation
)
