//Display / UI module for the emulator
package display

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var offset_x, offset_y float64

//Y value first then X
var Display [][]bool
var imd *imdraw.IMDraw

//64 X 32 PIXELS
//THUS 8 X 4 BYTES
const DISPLAY_WIDTH = 64
const DISPLAY_HEIGHT = 32
const CANVAS_WIDTH = DISPLAY_WIDTH * 8
const CANVAS_HEIGHT = DISPLAY_HEIGHT * 8

func CreateWindow(off_x, off_y float64) *pixelgl.Window {
	offset_x = off_x
	offset_y = off_y
	imd = imdraw.New(nil)
	Display = make([][]bool, DISPLAY_HEIGHT)
	for i := range Display {
		Display[i] = make([]bool, DISPLAY_WIDTH)
	}

	if off_x == 0 && off_y == 0 {
		//Initalize a matrix representing the display

		cfg := pixelgl.WindowConfig{
			Title:  "GoChip [Chip-8 Emulator]",
			Bounds: pixel.R(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT),
			VSync:  true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		return win
	}
	return nil
}

//From topleft
func push_square(imdraw *imdraw.IMDraw, x, y, width float64) {
	imdraw.Push(pixel.V(offset_x+x, offset_y+CANVAS_HEIGHT-y-width))
	imdraw.Push(pixel.V(offset_x+x+width, offset_y+CANVAS_HEIGHT-y-width))
	imdraw.Push(pixel.V(offset_x+x+width, offset_y+CANVAS_HEIGHT-y))
	imdraw.Push(pixel.V(offset_x+x, offset_y+CANVAS_HEIGHT-y))
	imdraw.Polygon(0)
}

func Render(win *pixelgl.Window) {
	for i := range Display {
		for k := range Display[i] {
			if Display[i][k] {
				push_square(imd, (float64(k * 8)), (float64(i * 8)), 8)
			}
		}
	}
	imd.Draw(win)
}

//Set every bit to false
func ClearScreen() {
	for i := range Display {
		for k := range Display[i] {
			Display[i][k] = false
		}
	}
}

//X and Y are supplied as pixel locations
//Returns true if the pixel was already enabled
func Draw(x, y int) bool {
	orig := Display[y][x]
	Display[y][x] = true
	return orig
}
