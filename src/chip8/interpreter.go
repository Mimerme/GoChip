package chip8

import "fmt"

//Create an address between 0x000 - 0xFFF given 3 bytes
func create_address(dig1 byte, dig2 byte, dig3 byte) uint16 {
	var address uint16
	address = ((uint16)(dig1)) << 8
	temp := ((uint16)(dig2)) << 4
	address = (address | temp) | ((uint16)(dig3))
	return address
}

func parse_opcode(high byte, low byte, machine *Chip8) {
	//First nibble (hex digit) of the opcode can POSSIBLY indicate the opcode [this is stored in opcode_id1]
	//If there are multiple opcodes that use the same nibble id then use the last nibble (last hex digit) [sotred in opcode_id2]
	var opcode_nib_1, opcode_nib_2, opcode_nib_3, opcode_nib_4 byte
	opcode_nib_1 = (high & 0xF0) >> 4
	opcode_nib_2 = (high & 0x0F)
	opcode_nib_3 = (low & 0xF0) >> 4
	opcode_nib_4 = (low & 0x0F)

	if opcode_nib_1 == 0x0 {
		if opcode_nib_4 == 0x0 {
			machine.display_clear()
		} else {
			machine.sub_ret()
		}
	} else if opcode_nib_1 == 0x1 {
		machine.jump(create_address(opcode_nib_2, opcode_nib_3, opcode_nib_4))
	} else if opcode_nib_1 == 0x2 {
		machine.call(create_address(opcode_nib_2, opcode_nib_3, opcode_nib_4))
	} else {
		fmt.Printf("Unknown opcode 0x" + string(opcode_nib_1) + string(opcode_nib_2) + string(opcode_nib_3) + string(opcode_nib_4))
	}
}
