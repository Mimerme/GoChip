package debugger

import (
	"fmt"
	"github.com/faiface/pixel/text"
)

func draw_registers_gl(txt *text.Text, registers [16]uint8, SP byte, PC uint16, I uint16) {
	fmt.Fprintln(txt, "Registers")
	for i, element := range registers {
		fmt.Fprintln(txt, "V[", i, "] |", to_hex_string_8(element))
	}
	fmt.Fprintln(txt, "I |", to_hex_string(I))
	fmt.Fprintln(txt, "PC |", to_hex_string(PC))
	fmt.Fprintln(txt, "SP |", to_hex_string_8(SP))
}

func draw_stack_gl(txt *text.Text, stack [16]uint16) {
	fmt.Fprintln(txt, "Stack")
	for i, element := range stack {
		fmt.Fprintln(txt, "Stack[", i, "] |", to_hex_string(element))
	}
}

func draw_memory_gl(txt *text.Text, memory [4096]byte) {

}
