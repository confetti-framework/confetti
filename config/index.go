package config

// Index of all configuration. Add your configuration to the list. This allows
// the framework and modules to use your configuration.
var Index = map[string]interface{}{
	"App":      App,
	"Path":     Path,
	"Errors":   Errors,
	"Logging":  Logging,
	"Embed":    Embed,
	"Database": Database,
}
