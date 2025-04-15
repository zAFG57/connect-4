package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func getCursorPosition(win *pixelgl.Window) (int,int) {
	return int(win.MousePosition().X), 280- int(win.MousePosition().Y)
}