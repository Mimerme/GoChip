package debugger

import (
	"fmt"
	"github.com/faiface/pixel/text"
)

func draw_registers_gl(txt *text.Text, registers [16]uint8, SP byte, PC, I uint16, DataTimer, SoundTimer byte) {
	fmt.Fprintln(txt, "Registers")
	for i, element := range registers {
		fmt.Fprintln(txt, "V[", i, "] |", to_hex_string_8(element))
	}
	fmt.Fprintln(txt, "I |", to_hex_string(I))
	fmt.Fprintln(txt, "PC |", to_hex_string(PC))
	fmt.Fprintln(txt, "SP |", to_hex_string_8(SP))
	fmt.Fprintln(txt, "DT |", (DataTimer))
	fmt.Fprintln(txt, "ST |", to_hex_string_8(SoundTimer))
}

func draw_stack_gl(txt *text.Text, stack [16]uint16) {
	fmt.Fprintln(txt, "Stack")
	for i, element := range stack {
		fmt.Fprintln(txt, "Stack[", i, "] |", to_hex_string(element))
	}
}

func draw_keys_gl(txt *text.Text, keys [16]byte) {
	fmt.Fprintln(txt, "Keyboard")
	for i, element := range keys {
		fmt.Fprintln(txt, "Key[", i, "] | ", element)
	}
}

func draw_memory_gl(txt *text.Text, memory [4096]byte) {

}
