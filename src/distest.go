package main

import "./display"
import "github.com/faiface/pixel/pixelgl"

func main() {
	pixelgl.Run(display.CreateWindow)
}
