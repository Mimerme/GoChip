//Display / UI module for the emulator
package display

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//Y value first then X
var Display [][]bool
var imd *imdraw.IMDraw
var win *pixelgl.Window
var Ready bool = false

//64 X 32 PIXELS
//THUS 8 X 4 BYTES
const DISPLAY_WIDTH = 64
const DISPLAY_HEIGHT = 32
const CANVAS_WIDTH = DISPLAY_WIDTH * 8
const CANVAS_HEIGHT = DISPLAY_HEIGHT * 8

func CreateWindow() {
	//Initalize a matrix representing the display
	Display = make([][]bool, DISPLAY_HEIGHT)
	for i := range Display {
		Display[i] = make([]bool, DISPLAY_WIDTH)
	}

	cfg := pixelgl.WindowConfig{
		Title:  "GoChip [Chip-8 Emulator]",
		Bounds: pixel.R(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT),
		VSync:  true,
	}
	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd = imdraw.New(nil)
	for !win.Closed() {
		Render()
	}

	Ready = true
}

//From topleft
func push_square(imdraw *imdraw.IMDraw, x, y, width float64) {
	imdraw.Push(pixel.V(x, CANVAS_HEIGHT-y-width))
	imdraw.Push(pixel.V(x+width, CANVAS_HEIGHT-y-width))
	imdraw.Push(pixel.V(x+width, CANVAS_HEIGHT-y))
	imdraw.Push(pixel.V(x, CANVAS_HEIGHT-y))
	imdraw.Polygon(0)
}

func Render() {
	for i := range Display {
		for k := range Display[i] {
			if Display[i][k] {
				push_square(imd, (float64(k * 8)), (float64(i * 8)), 8)
			}
		}
	}
	win.Clear(colornames.Black)
	imd.Draw(win)
	win.Update()
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
	fmt.Println(Display)
	orig := Display[y][x]
	Display[y][x] = true
	return orig
}
