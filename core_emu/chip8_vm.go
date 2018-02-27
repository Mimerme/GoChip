package main

const STACK_START uint16 = 0x0EA0
const STACK_END uint16 = 0x0EFF
const DISP_START uint16 = 0x0F00
const DISP_END uint16 = 0x0FFF

//Contains the machine specification structure
type Chip8 struct {
	//Allocate the 4096 bytes of memory
	//TODO: Implement invalid program memory access
	//https://en.wikipedia.org/wiki/CHIP-8#Memory
	memory [4096]byte

	//Registers
	//Register 15 (aka 16) is VF
	GPR [16]byte
	I   uint16
	//PC is only 3 nibbles
	PC uint16

	//The stack has a max depth of 16
	SP byte
}
