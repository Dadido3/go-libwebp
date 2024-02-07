package webp_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/Dadido3/go-libwebp/test/util"
	"github.com/Dadido3/go-libwebp/webp"
)

func BenchmarkEncodeRGBA(b *testing.B) {

	img := util.ReadPNG(filepath.Join(".", "..", "examples", "images", "yellow-rose-3.png"))

	config, err := webp.ConfigPreset(webp.PresetDefault, 100)
	if err != nil {
		b.Fatalf("ConfigPreset failed: %v", err)
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 1000000))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := webp.EncodeRGBA(buffer, img, config); err != nil {
			b.Errorf("EncodeRGBA failed: %v", err)
			return
		}

		buffer.Reset()
	}
}

func BenchmarkEncodeYUVA(b *testing.B) {

	data := util.ReadFile(filepath.Join(".", "..", "examples", "images", "cosmos.webp"))
	options := &webp.DecoderOptions{}

	img, err := webp.DecodeYUVA(data, options)
	if err != nil {
		b.Errorf("DecodeYUVA failed: %v", err)
		return
	}

	config, err := webp.ConfigPreset(webp.PresetDefault, 100)
	if err != nil {
		b.Fatalf("ConfigPreset failed: %v", err)
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 1000000))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := webp.EncodeYUVA(buffer, img, config); err != nil {
			b.Errorf("EncodeYUVA failed: %v", err)
			return
		}

		buffer.Reset()
	}
}

func BenchmarkEncodeLossless9(b *testing.B) {

	img := util.ReadPNG(filepath.Join(".", "..", "examples", "images", "fizyplankton.png"))

	config, err := webp.ConfigLosslessPreset(9)
	if err != nil {
		b.Fatalf("ConfigPreset failed: %v", err)
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 10000000))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := webp.Encode(buffer, img, config); err != nil {
			b.Errorf("Encode failed: %v", err)
			return
		}

		buffer.Reset()
	}
}

func BenchmarkEncodeLossless8(b *testing.B) {

	img := util.ReadPNG(filepath.Join(".", "..", "examples", "images", "fizyplankton.png"))

	config, err := webp.ConfigLosslessPreset(8)
	if err != nil {
		b.Fatalf("ConfigPreset failed: %v", err)
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 10000000))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := webp.Encode(buffer, img, config); err != nil {
			b.Errorf("Encode failed: %v", err)
			return
		}

		buffer.Reset()
	}
}

func BenchmarkEncodeLossless7(b *testing.B) {

	img := util.ReadPNG(filepath.Join(".", "..", "examples", "images", "fizyplankton.png"))

	config, err := webp.ConfigLosslessPreset(7)
	if err != nil {
		b.Fatalf("ConfigPreset failed: %v", err)
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 10000000))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := webp.Encode(buffer, img, config); err != nil {
			b.Errorf("Encode failed: %v", err)
			return
		}

		buffer.Reset()
	}
}
