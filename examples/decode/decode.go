// Package main is an example implementation of WebP decoder.
package main

import (
	"path/filepath"

	"github.com/Dadido3/go-libwebp/test/util"
	"github.com/Dadido3/go-libwebp/webp"
)

func main() {
	var err error

	// Read binary data
	data := util.ReadFile(filepath.Join(".", "..", "images", "cosmos.webp"))

	// Decode
	options := &webp.DecoderOptions{}
	img, err := webp.DecodeRGBA(data, options)
	if err != nil {
		panic(err)
	}

	util.WritePNG(img, filepath.Join(".", "..", "out", "encoded_cosmos.png"))
}
