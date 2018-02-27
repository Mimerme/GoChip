package main

import "fmt"

var machine Chip8

func Initialize() {
	machine = Chip8{}

	for i := 0; i < len(machine.memory); i++ {
		machine.memory[i] = 0x00
	}
	for i := 0; i < len(machine.GPR); i++ {
		machine.GPR[i] = 0x00
	}
	machine.I = 0x0000
	machine.PC = 0x0000
}

//Create an address between 0x000 - 0xFFF given 3 bytes
func create_address(dig1 byte, dig2 byte, dig3 byte) uint16 {
	var address uint16
	address = ((uint16)(dig1)) << 8
	temp := ((uint16)(dig2)) << 4
	address = (address & temp) & ((uint16)(dig3))
	return address
}

func parse_opcode(high byte, low byte) {
	//First nibble (hex digit) of the opcode can POSSIBLY indicate the opcode [this is stored in opcode_id1]
	//If there are multiple opcodes that use the same nibble id then use the last nibble (last hex digit) [sotred in opcode_id2]
	var opcode_nib_1, opcode_nib_2, opcode_nib_3, opcode_nib_4 byte
	opcode_nib_1 = (high & 0xF0) >> 12
	opcode_nib_4 = low & 0x0F
	opcode_nib_2 = high & 0x0F
	opcode_nib_3 = (low & 0xF0) >> 12

	if opcode_nib_1 == 0x0 {
		if opcode_nib_4 == 0x0 {
			display_clear()
		} else {
			sub_ret()
		}
	} else if opcode_nib_1 == 0x1 {
		jump(create_address(opcode_nib_2, opcode_nib_3, opcode_nib_4))
	} else if opcode_nib_1 == 0x2 {
		call(create_address(opcode_nib_2, opcode_nib_3, opcode_nib_4))
	} else {
		fmt.Printf("Unknown opcode 0x" + string(opcode_nib_1) + string(opcode_nib_2) + string(opcode_nib_3) + string(opcode_nib_4))
	}
}
