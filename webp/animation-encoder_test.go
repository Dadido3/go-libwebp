package webp_test

import (
	"image"
	"path/filepath"
	"testing"
	"time"

	"github.com/Dadido3/go-libwebp/test/util"
	"github.com/Dadido3/go-libwebp/webp"
)

func TestEncodeAnimation(t *testing.T) {
	data := util.ReadFile(filepath.Join(".", "..", "examples", "images", "cosmos.webp"))
	aWebP, err := webp.DecodeRGBA(data, &webp.DecoderOptions{})
	if err != nil {
		t.Fatalf("Got Error: %v", err)
	}

	img := []image.Image{
		util.ReadPNG(filepath.Join(".", "..", "examples", "images", "butterfly.png")),
		util.ReadPNG(filepath.Join(".", "..", "examples", "images", "checkerboard.png")),
		util.ReadPNG(filepath.Join(".", "..", "examples", "images", "yellow-rose-3.png")),
		aWebP,
	}

	width, height := 24, 24
	anim, err := webp.NewAnimationEncoder(width, height, 0, 0)
	if err != nil {
		t.Fatalf("initializing decoder: %v", err)
	}
	defer anim.Close()

	for i, im := range img {
		// all frames of an animation must have the same dimensions
		cropped := im.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(0, 0, width, height))

		if err := anim.AddFrame(cropped, 100*time.Millisecond); err != nil {
			t.Errorf("adding frame %d: %v", i, err)
		}
	}

	buf, err := anim.Assemble()
	if err != nil {
		t.Fatalf("assembling animation: %v", err)
	}
	if len(buf) == 0 {
		t.Errorf("assembled animation is empty")
	}
}
