package data

import "embed"

//go:embed assets
var Asset embed.FS

//go:embed embed
var Embed embed.FS
