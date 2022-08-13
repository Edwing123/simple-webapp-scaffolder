package resources

import "embed"

//go:embed  resources/views
var ViewsFS embed.FS

//go:embed  resources/assets
var AssetsFS embed.FS

//go:embed  all:resources/configs
var ConfigsFS embed.FS
