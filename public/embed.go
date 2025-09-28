package public

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "dist")
}

func Manifest() ([]byte, error) {
	return assets.ReadFile("dist/.vite/manifest.json")
}
