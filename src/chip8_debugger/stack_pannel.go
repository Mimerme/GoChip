package chip8_debugger

import "github.com/marcusolsson/tui-go"
import (
	"fmt"
	"image"
	"strconv"
)

func draw_stack(draw_pane *tui.Box, stack *[16]uint16) {
	//4 bytes on each row
	//2 bytes for address
	//2 bytes for actually return addres after subroutine
	//16 rows of data
	grid := tui.NewGrid(2, 17)
	grid.SetBorder(true)

	address_label := tui.NewLabel("Stack Pointer")
	address_label.SetSizePolicy(tui.Minimum, tui.Minimum)

	value_label := tui.NewLabel("Value")
	value_label.SetSizePolicy(tui.Minimum, tui.Minimum)

	grid.SetCell(image.Point{X: 0, Y: 0}, address_label)
	grid.SetCell(image.Point{X: 1, Y: 0}, value_label)

	//Read values
	for i := 0; i < 16; i++ {
		grid.SetCell(image.Point{X: 0, Y: i + 1}, tui.NewLabel(strconv.Itoa(i)))
		//TODO: Implement reading the stack
	}

	grid.SetSizePolicy(tui.Minimum, tui.Minimum)
	draw_pane.Append(grid)
}

func draw_registers(draw_pane *tui.Box, GPRs *[16]byte, I *uint16, PC *uint16, SP *byte) {
	grid := tui.NewGrid(4, 16)
	grid.SetBorder(true)

	for i := 0; i < 16; i++ {
		grid.SetCell(image.Point{X: 0, Y: i}, tui.NewLabel(fmt.Sprintf("V[%X]", i)))
		//TODO: Implement reading the stack
	}
	pc_label := tui.NewLabel("0")

	go func(label *tui.Label) {
		for {
			h := fmt.Sprintf("%x", PC)
			pc_label.SetText(h)
		}
	}(pc_label)

	grid.SetCell(image.Point{X: 2, Y: 0}, tui.NewLabel("I"))
	grid.SetCell(image.Point{X: 2, Y: 1}, tui.NewLabel("PC"))
	grid.SetCell(image.Point{X: 3, Y: 1}, pc_label)
	grid.SetCell(image.Point{X: 2, Y: 2}, tui.NewLabel("SP"))

	grid.SetSizePolicy(tui.Minimum, tui.Minimum)
	draw_pane.Append(grid)
}
