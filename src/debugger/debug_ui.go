package debugger

import "golang.org/x/image/font/basicfont"

import "fmt"
import "../chip8/"
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

var current_pane = 1
var basicAtlas *text.Atlas
var basicTxt, basicTxt2, basicTxt3 *text.Text
var footer *text.Text

const WIDTH = 700 + 64
const HEIGHT = 400

func CreateWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "GoChip Debugger",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//	face, err := loadTTF("intuitive.ttf", 52)
	//	if err != nil {
	//		panic(err)
	//	}
	//

	//Initalize the text rendering requirenments
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt = text.New(pixel.V(0, HEIGHT-10), basicAtlas)
	basicTxt2 = text.New(pixel.V(150, HEIGHT-10), basicAtlas)
	basicTxt3 = text.New(pixel.V(300, HEIGHT-10), basicAtlas)
	footer = text.New(pixel.V(0, 10), basicAtlas)
	fmt.Fprintln(footer, "<Tab> to cycle through pannels. <P> to pause execution. <D> to step")
	return win
}

func Render(win *pixelgl.Window, machine *chip8.Chip8, paused, execute_next *bool) {
	if win.JustPressed(pixelgl.KeyP) {
		fmt.Println("Toggling execution...")
		(*paused) = !(*paused)
		footer.Clear()
		if *paused {
			fmt.Fprintln(footer, "Execution paused...")
		} else {
			fmt.Fprintln(footer, "Execution resumed")
		}
	}

	if win.JustPressed(pixelgl.KeyD) && (*paused) {
		fmt.Println("Executing next instruction")
		footer.Clear()
		fmt.Fprintln(footer, "Executed next step")
		(*execute_next) = true
	}

	basicTxt.Clear()
	basicTxt2.Clear()
	basicTxt3.Clear()

	draw_registers_gl(basicTxt, machine.GPR, machine.SP, machine.PC, machine.I, machine.DT, machine.ST)
	draw_stack_gl(basicTxt2, machine.Stack)
	draw_keys_gl(basicTxt3, machine.Keys)

	basicTxt.Draw(win, pixel.IM)
	basicTxt2.Draw(win, pixel.IM)
	basicTxt3.Draw(win, pixel.IM)
	footer.Draw(win, pixel.IM)
}
