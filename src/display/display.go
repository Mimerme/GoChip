//Display / UI module for the emulator
package display

import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

func CreateWindow() {
	cfg := pixelgl.WindowConfig{
		Title:  "GoChip [Chip-8 Emulator]",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	_, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
}

func ClearScreen() {

}

func Draw(x, y int) {

}
