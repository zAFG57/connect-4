package main

import (
	"image"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func createWindow(img image.Image) (*pixelgl.Window, error) {

	cfg := pixelgl.WindowConfig{
		Title:  "p4",
		Bounds: pixel.R(0, 0, float64(img.Bounds().Max.X), float64(img.Bounds().Max.Y)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}

	return win, nil
}

func UpdateWindow(win *pixelgl.Window, img *image.RGBA) {
	pic := pixel.PictureDataFromImage(img)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(pixel.RGB(0, 0, 0))
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	win.Update()
}