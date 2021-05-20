package embed

import (
	"embed"
	"io/fs"
)

//go:embed data/*
var assets embed.FS

func Assets() fs.FS {
	assetsFs, err := fs.Sub(assets, "data")
	if err != nil {
		panic(err)
	}
	return assetsFs
}
