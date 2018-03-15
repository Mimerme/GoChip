package main

import "./display"
import "github.com/faiface/pixel/pixelgl"
import "fmt"

func main() {
	go func() {
		for !display.Ready {
			fmt.Println("not ready")
		}
		fmt.Println("ready!")
	}()

	pixelgl.Run(display.CreateWindow)
}
