// Package util contains utility code for demonstration of go-libwebp.
package util

import (
	"bufio"
	"image"
	"image/png"
	"io"
	"os"
)

// OpenFile opens specified example file
func OpenFile(filePath string) (io io.Reader) {
	io, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return
}

// ReadFile reads and returns data bytes of specified example file.
func ReadFile(filePath string) (data []byte) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return
}

// CreateFile opens specified example file
func CreateFile(filePath string) (f *os.File) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return
}

// WritePNG encodes and writes image into PNG file.
func WritePNG(img image.Image, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(f)
	defer func() {
		b.Flush()
		f.Close()
	}()

	if err := png.Encode(b, img); err != nil {
		panic(err)
	}
}

// ReadPNG reads and decodes png data into image.Image
func ReadPNG(filePath string) (img image.Image) {
	io, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	img, err = png.Decode(io)
	if err != nil {
		panic(err)
	}
	return
}
