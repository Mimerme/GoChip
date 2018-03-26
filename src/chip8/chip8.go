package chip8

import "time"

func InitializeVM() *Chip8 {
	machine := Chip8{}
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
