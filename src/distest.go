package main

import "./display"
import "github.com/faiface/pixel/pixelgl"
import "os"

func main() {
	pixelgl.Run(
		func() {
			display.CreateWindow()
		},
	)
}
