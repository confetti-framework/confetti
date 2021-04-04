package views

import "embed"

// Views provides a collection of all views and templates. This is to load all
// predefined templates.
//go:embed *
var Views embed.FS
