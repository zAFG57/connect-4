package main

import (
	"image"
	"time"
	"github.com/faiface/pixel/pixelgl"
)

var isGameRuning bool = true
var isTimedOut bool = false
var game Game
var p1 Player = &HumanPlayer{}
var p2 Player = &HumanPlayer{}
var p3 *IaPlayer = &IaPlayer{}

func main() {
	pixelgl.Run(run)
}

func run() {
	img := image.NewRGBA(image.Rect(0, 0, 280, 280))
	win, _ := createWindow(img)
	defer win.Destroy()
	p1.Init(&game)
	p2.Init(&game)
	p3.Init(&game)
	game.Init(p1, p3, img)
	UpdateWindow(win,img)
	playingLoop(win,img)
}

func playingLoop(win *pixelgl.Window, img *image.RGBA) {
	for ;!win.Closed(); {
		if !isTimedOut && win.Pressed(pixelgl.MouseButtonLeft) && isGameRuning {
			x,_ := getCursorPosition(win)
			game.Click(x)
			go func() {
				isTimedOut = true
				time.Sleep(200 * time.Millisecond)
				isTimedOut = false
			}()
		}
		UpdateWindow(win,img)
	}
}