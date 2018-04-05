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
	switch nib_1 {
	case 0x0:
		if nib_2 == 0 && nib_3 == 0xE {
			if nib_4 == 0x0 {
				machine.CLS()
			} else if nib_4 == 0xE {
				machine.RET()
			}
		}
	case 0x1:
		machine.JP(((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4))
	case 0x2:
		machine.CALL(((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4))
	case 0x3:
		machine.SE(nib_2, ((nib_3 << 4) | (nib_4)))
	case 0x4:
		machine.SNE(nib_2, ((nib_3 << 4) | (nib_4)))
	case 0x5:
		if nib_4 == 0 {
			machine.SE_REG(nib_2, nib_3)
		}
	case 0x6:
		machine.LD(nib_2, nib_3, nib_4)
	case 0x7:
		machine.ADD(nib_2, nib_3, nib_4)
	case 0x8:
		switch nib_4 {
		case 0x0:
			machine.LD_REG(nib_2, nib_3)
		case 0x1:
			machine.OR(nib_2, nib_3)
		case 0x2:
			machine.AND(nib_2, nib_3)
		case 0x3:
			machine.XOR(nib_2, nib_3)
		case 0x4:
			machine.ADD_REG(nib_2, nib_3)
		case 0x5:
			machine.SUB(nib_2, nib_3)
		case 0x6:
			machine.SHR(nib_2, nib_3)
		case 0x7:
			machine.SUBN(nib_2, nib_3)
		case 0xE:
			machine.SHL(nib_2, nib_3)
		default:
			fmt.Printf("Unknown opcode 0x", string(nib_1), string(nib_2), string(nib_3), string(nib_4))
		}
	case 0x9:
		if nib_4 == 0 {
			machine.SNE_REG(nib_2, nib_3)
		}
	case 0xA:
		machine.LD_I(nib_2, nib_3, nib_4)
	case 0xB:
		machine.JP_V0(nib_2, nib_3, nib_4)
	case 0xC:
		machine.RND((nib_3<<4)|(nib_4), nib_2)
	case 0xD:
		machine.DRW(nib_4, machine.GPR[nib_2], machine.GPR[nib_3])
	case 0xE:
		if nib_3 == 0x9 && nib_4 == 0xE {
			machine.SKP(nib_2)
		} else if nib_3 == 0xA && nib_4 == 0x1 {
			machine.SKNP(nib_2)
		}
	case 0xF:
		if nib_3 == 0x0 && nib_4 == 0xA {
			machine.BLOCK_KEY(nib_2)
		} else if nib_3 == 0x0 && nib_4 == 0x7 {
			machine.STR_DT(nib_2)
		} else if nib_3 == 0x1 && nib_4 == 0x5 {
			machine.LD_DT(nib_2)
		} else if nib_3 == 0x1 && nib_4 == 0x8 {
			machine.LD_ST(nib_2)
		} else if nib_3 == 0x1 && nib_4 == 0xE {
			machine.ADD_I(nib_2)
		} else if nib_3 == 0x2 && nib_4 == 0x9 {
			machine.LD_TXT_SPRITE(nib_2)
		} else if nib_3 == 0x3 && nib_4 == 0x3 {
			machine.STR_BCD(nib_2)
		} else if nib_3 == 0x5 && nib_4 == 0x5 {
			machine.STR_MULTI(nib_2)
		} else if nib_3 == 0x6 && nib_4 == 0x5 {
			machine.LD_MULTI(nib_2)
		}

	default:
		fmt.Printf("Unknown opcode 0x", string(nib_1), string(nib_2), string(nib_3), string(nib_4))
	}
}
