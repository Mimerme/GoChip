package debugger

import "golang.org/x/image/font/basicfont"

import "github.com/marcusolsson/tui-go"
import "fmt"
import "../chip8/"
import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)
import (
	"golang.org/x/image/font"
	"io/ioutil"
	"os"
)
import "github.com/golang/freetype/truetype"

var current_pane = 1
var basicAtlas *text.Atlas
var basicTxt *text.Text
var basicTxt2 *text.Text
var footer *text.Text

const WIDTH = 700 + 64
const HEIGHT = 400

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

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

	switch current_pane {
	case 1:
		basicTxt.Clear()
		basicTxt2.Clear()
		draw_registers_gl(basicTxt, machine.GPR, machine.SP, machine.PC, machine.I)
		draw_stack_gl(basicTxt2, machine.Stack)
		break
	}
	basicTxt.Draw(win, pixel.IM)
	basicTxt2.Draw(win, pixel.IM)
	footer.Draw(win, pixel.IM)
}

func StartDebugger(chip8VM *(chip8.Chip8), DEBUG_PAUSE *bool, pause *chan struct{}, play *chan struct{}, step *chan struct{}) {
	fmt.Println("Starting the debugger")

	reg_label := tui.NewVBox(
		tui.NewLabel("Registers"),
	)
	reg_label.SetSizePolicy(tui.Minimum, tui.Maximum)

	stack_label := tui.NewVBox(
		tui.NewLabel("Stack"),
	)
	stack_label.SetSizePolicy(tui.Minimum, tui.Maximum)

	footer := tui.NewVBox(
		tui.NewLabel("<Tab> to cycle through panels. <P> to pause execution. <D> Execute instruction"),
	)
	footer.SetSizePolicy(tui.Minimum, tui.Maximum)

	registers := tui.NewVBox(reg_label)
	stack := tui.NewVBox(stack_label)
	footer_box := tui.NewVBox(footer)

	status_bar := tui.NewStatusBar("Registers & Stack")
	status_bar.SetPermanentText("Executing Program...")

	main_panel := tui.NewVBox(
		status_bar,
		tui.NewHBox(registers, stack),
		footer_box,
	)

	registers.SetBorder(true)
	stack.SetBorder(true)

	//Draw the stack
	draw_stack(stack, &chip8VM.Stack)
	draw_registers(registers, &chip8VM.GPR, &chip8VM.I, &chip8VM.PC, &chip8VM.SP)

	ui, err := tui.New(main_panel)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	ui.SetKeybinding("P", func() {
		*(DEBUG_PAUSE) = !(*(DEBUG_PAUSE))
		if *DEBUG_PAUSE {
			status_bar.SetPermanentText("Execution Paused")
			(*pause) <- struct{}{}
		} else {
			status_bar.SetPermanentText("Executing Program...")
			(*play) <- struct{}{}
		}
	})

	ui.SetKeybinding("D", func() {
		//Run a step
		if *DEBUG_PAUSE {
			status_bar.SetPermanentText("Stepped over instruction")
			(*chip8VM).ExecuteStep()
		}
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
