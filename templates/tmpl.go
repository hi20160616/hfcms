package templates

import "embed"

//go:embed bootstrap default/*.html
var FS embed.FS
