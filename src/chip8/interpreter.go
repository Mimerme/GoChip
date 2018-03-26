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
	var nib_1, nib_2, nib_3, nib_4 byte
	nib_1 = (high & 0xF0) >> 4
	nib_2 = (high & 0x0F)
	nib_3 = (low & 0xF0) >> 4
	nib_4 = (low & 0x0F)

	//TODO Refine this parsing
	if nib_1 == 0x0 {
		if nib_4 == 0x0 && nib_3 == 0xE {
			machine.display_clear()
			machine.PC += 2
		} else if nib_4 == 0xE && nib_3 == 0xE {
			machine.sub_ret()
		}
	} else if nib_1 == 0x1 {
		machine.jump(create_address(nib_2, nib_3, nib_4))
	} else if nib_1 == 0x2 {
		machine.call(create_address(nib_2, nib_3, nib_4))
	} else if nib_1 == 0x3 {
		machine.skip_if_equal(nib_2, ((nib_3 << 4) | nib_4))
	} else if nib_1 == 0x4 {
		machine.skip_if_not_equal(nib_2, ((nib_3 << 4) | nib_4))
	} else if nib_1 == 0x5 && nib_4 == 0x0 {
		machine.skip_if_equal_reg(nib_2, nib_3)
	} else if nib_1 == 0x6 {
		machine.GPR[nib_2] = (nib_3 << 4) | nib_4
		machine.PC += 2
	} else if nib_1 == 0x7 {
		machine.GPR[nib_2] = machine.GPR[nib_2] + ((nib_3 << 4) | nib_4)
		machine.PC += 2
	} else if nib_1 == 0x8 {
		if nib_4 == 0x1 {
			machine.GPR[nib_2] = (machine.GPR[nib_3]) | machine.GPR[nib_2]
			machine.PC += 2
		} else if nib_4 == 0x2 {
			machine.GPR[nib_2] = (machine.GPR[nib_3]) & machine.GPR[nib_2]
			machine.PC += 2
		}
	} else if nib_1 == 0x9 && nib_4 == 0x0 {
		machine.skip_if_not_equal_reg(nib_2, nib_3)
	} else if nib_1 == 0xA {
		machine.I = ((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4)
		machine.PC += 2
	} else if nib_1 == 0xB {
		machine.PC = ((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4)
		machine.PC += (uint16)(machine.GPR[0])
	} else if nib_1 == 0xC {
		//TODO: Generate random
		machine.generate_random((nib_3<<4)|nib_4, nib_2)
		machine.PC += 2
	} else if nib_1 == 0xD {
		//TODO: Draw sprite
		machine.draw_sprite(nib_4, nib_2, nib_3)
		machine.PC += 2
	} else if nib_1 == 0xE && nib_3 == 0x9 && nib_4 == 0xE {
		//TODO: Skip if key
		machine.PC += 2
		if machine.Keys[nib_2] == 1 {
			machine.PC += 2
		}
	} else if nib_1 == 0xE && nib_3 == 0xA && nib_4 == 0x1 {
		//TODO: skip if not key
		machine.PC += 2
		if machine.Keys[nib_2] != 1 {
			machine.PC += 2
		}
	} else if nib_1 == 0xF {
		if nib_3 == 0x0 && nib_4 == 0x7 {
			//TODO: Set register to delay timer
			machine.GPR[nib_2] = machine.DT
			machine.PC += 2
		} else if nib_3 == 0x0 && nib_4 == 0xA {
			//TODO: block and wait for key press
		} else if nib_3 == 0x1 && nib_4 == 0x5 {
			//TODO: delay timer to reg
			machine.DT = machine.GPR[nib_2]
			machine.PC += 2
		} else if nib_3 == 0x1 && nib_4 == 0x8 {
			//TODO: sound timer set
			machine.ST = machine.GPR[nib_2]
			machine.PC += 2
		} else if nib_3 == 0x1 && nib_4 == 0xE {
			//machine.I += machine.GPR[nib_2]
		} else if nib_3 == 0x2 && nib_4 == 0x9 {
		} else if nib_3 == 0x3 && nib_4 == 0x3 {

		} else if nib_3 == 0x5 && nib_4 == 0x5 {
		} else if nib_3 == 0x6 && nib_4 == 0x5 {
		}
	} else {
		fmt.Printf("Unknown opcode 0x", string(nib_1), string(nib_2), string(nib_3), string(nib_4))
	}
}
