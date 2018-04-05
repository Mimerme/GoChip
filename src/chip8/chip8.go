package chip8

import "time"

func InitializeVM() *Chip8 {
	machine := Chip8{}

	//Load generic text-sprites
	//0
	machine.Memory[0x0] = 0xF0
	machine.Memory[0x0+1] = 0x90
	machine.Memory[0x0+2] = 0x90
	machine.Memory[0x0+3] = 0x90
	machine.Memory[0x0+4] = 0xF0
	//1
	machine.Memory[0x5] = 0x20
	machine.Memory[0x5+1] = 0x60
	machine.Memory[0x5+2] = 0x20
	machine.Memory[0x5+3] = 0x20
	machine.Memory[0x5+4] = 0x70
	//2
	machine.Memory[0xA] = 0xF0
	machine.Memory[0xA+1] = 0x10
	machine.Memory[0xA+2] = 0xF0
	machine.Memory[0xA+3] = 0x80
	machine.Memory[0xA+4] = 0xF0
	//3
	machine.Memory[0xF] = 0xF0
	machine.Memory[0xF+1] = 0x10
	machine.Memory[0xF+2] = 0xF0
	machine.Memory[0xF+3] = 0x10
	machine.Memory[0xF+4] = 0xF0
	//4
	machine.Memory[0x14] = 0x90
	machine.Memory[0x14+1] = 0x90
	machine.Memory[0x14+2] = 0xF0
	machine.Memory[0x14+3] = 0x10
	machine.Memory[0x14+4] = 0x10
	//5
	machine.Memory[0x19] = 0xF0
	machine.Memory[0x19+1] = 0x80
	machine.Memory[0x19+2] = 0xF0
	machine.Memory[0x19+3] = 0x10
	machine.Memory[0x19+4] = 0xF0
	//6
	machine.Memory[0x1E] = 0xF0
	machine.Memory[0x1E+1] = 0x80
	machine.Memory[0x1E+2] = 0xF0
	machine.Memory[0x1E+3] = 0x90
	machine.Memory[0x1E+4] = 0xF0
	//7
	machine.Memory[0x23] = 0xF0
	machine.Memory[0x23+1] = 0x10
	machine.Memory[0x23+2] = 0x20
	machine.Memory[0x23+3] = 0x40
	machine.Memory[0x23+4] = 0x40
	//8
	machine.Memory[0x28] = 0xF0
	machine.Memory[0x28+1] = 0x90
	machine.Memory[0x28+2] = 0xF0
	machine.Memory[0x28+3] = 0x90
	machine.Memory[0x28+4] = 0xF0
	//9
	machine.Memory[0x2D] = 0xF0
	machine.Memory[0x2D+1] = 0x90
	machine.Memory[0x2D+2] = 0xF0
	machine.Memory[0x2D+3] = 0x10
	machine.Memory[0x2D+4] = 0xF0
	//A
	machine.Memory[0x32] = 0xF0
	machine.Memory[0x32+1] = 0x90
	machine.Memory[0x32+2] = 0xF0
	machine.Memory[0x32+3] = 0x90
	machine.Memory[0x32+4] = 0x90
	//B
	machine.Memory[0x37] = 0xE0
	machine.Memory[0x37+1] = 0x90
	machine.Memory[0x37+2] = 0xE0
	machine.Memory[0x37+3] = 0x90
	machine.Memory[0x37+4] = 0xE0
	//C
	machine.Memory[0x3C] = 0xF0
	machine.Memory[0x3C+1] = 0x80
	machine.Memory[0x3C+2] = 0x80
	machine.Memory[0x3C+3] = 0x80
	machine.Memory[0x3C+4] = 0xF0
	//D
	machine.Memory[0x41] = 0xE0
	machine.Memory[0x41+1] = 0x90
	machine.Memory[0x41+2] = 0x90
	machine.Memory[0x41+3] = 0x90
	machine.Memory[0x41+4] = 0xE0
	//E
	machine.Memory[0x46] = 0xF0
	machine.Memory[0x46+1] = 0x80
	machine.Memory[0x46+2] = 0xF0
	machine.Memory[0x46+3] = 0x80
	machine.Memory[0x46+4] = 0xF0
	//F
	machine.Memory[0x4B] = 0xF0
	machine.Memory[0x4B+1] = 0x80
	machine.Memory[0x4B+2] = 0xF0
	machine.Memory[0x4B+3] = 0x80
	machine.Memory[0x4B+4] = 0x80

	return &machine
}

const DISP_START uint16 = 0x0F00
const DISP_END uint16 = 0x0FFF

//Contains the machine specification structure
type Chip8 struct {
	//Allocate the 4096 bytes of memory
	//TODO: Implement invalid program memory access
	//https://en.wikipedia.org/wiki/CHIP-8#Memory
	Memory [4096]byte

	//Registers
	//Register 15 (aka 16) is VF
	GPR [16]byte
	I   uint16
	//PC is only 3 nibbles
	PC uint16

	//The stack has a max depth of 16
	SP    byte
	Stack [16]uint16

	Keys [16]byte

	//Delay timer
	DT byte
	//Sound timer
	ST byte
}

func StartDelayTimer(machine *Chip8) {
	for {
		if machine.DT != 0 {
			machine.DT -= 1
		}
		//Should be 16.6 for 60Hz, but time.Sleep requires integers
		time.Sleep((17) * time.Millisecond)
	}
}

func StartSoundTimer(machine *Chip8) {
	for {
		if machine.ST != 0 {
			machine.ST -= 1
		}
		//Should be 16.6 for 60Hz, but time.Sleep requires integers
		time.Sleep((17) * time.Millisecond)
	}
}
