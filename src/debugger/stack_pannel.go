package debugger

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
	pc_label := tui.NewLabel("NoInit")
	i_label := tui.NewLabel("NoInit")
	sp_label := tui.NewLabel("NoInit")

	hook_update_label(pc_label, PC)
	hook_update_label(i_label, I)
	hook_update_label_8(sp_label, SP)

	grid.SetCell(image.Point{X: 2, Y: 0}, tui.NewLabel("I"))
	grid.SetCell(image.Point{X: 3, Y: 0}, i_label)
	grid.SetCell(image.Point{X: 2, Y: 1}, tui.NewLabel("PC"))
	grid.SetCell(image.Point{X: 3, Y: 1}, pc_label)
	grid.SetCell(image.Point{X: 2, Y: 2}, tui.NewLabel("SP"))
	grid.SetCell(image.Point{X: 3, Y: 2}, sp_label)

	grid.SetSizePolicy(tui.Minimum, tui.Minimum)
	draw_pane.Append(grid)
}

//Given a label set the value that it should watch
func hook_update_label(label *tui.Label, val *uint16) {
	go func() {
		for {
			label.SetText("0x" + to_hex_string(*val))
		}
	}()
}

func to_hex_string(val uint16) (hex_string string) {
	return strconv.FormatUint((uint64)(val), 16)
}

//Same as funcs above just with 8 bit unsigned ints
func hook_update_label_8(label *tui.Label, val *uint8) {
	go func() {
		for {
			label.SetText("0x" + to_hex_string_8(*val))
		}
	}()
}

func to_hex_string_8(val uint8) (hex_string string) {
	return strconv.FormatUint((uint64)(val), 8)
}
