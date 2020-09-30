module lanvard

go 1.15

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/foundation v0.2.2
	github.com/lanvard/routing v0.2.0
	github.com/lanvard/support v0.1.0
	github.com/stretchr/testify v1.5.1
	golang.org/x/text v0.3.2
)

replace (
	github.com/lanvard/contract v0.0.0 => ../github.com/lanvard/contract
	github.com/lanvard/foundation v0.2.2 => ../github.com/lanvard/foundation
	github.com/lanvard/routing v0.2.0 => ../github.com/lanvard/routing
	github.com/lanvard/support v0.1.0 => ../github.com/lanvard/support
)
