package data

import "embed"

//go:embed embed
var Embed embed.FS

//go:embed model
var Model embed.FS

//go:embed tpl
var TPL embed.FS
