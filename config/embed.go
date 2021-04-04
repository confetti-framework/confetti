package config

import (
	"embed"
	"src/resources/views"
)

// Embed contains fields of all resources to be loaded during compile time.
var Embed = struct {
	Views embed.FS
}{
	Views: views.Views,
}
