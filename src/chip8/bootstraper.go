package chip8

func (machine *Chip8) BootstrapProgram(opcodes []OpCode) {
	for i := 0; i < len(opcodes); i++ {
		//Load the code into memory
		//Each instruction is 2 bytes
		machine.Memory[0x200+(i*2)] = opcodes[i].high
		machine.Memory[0x200+(i*2)+1] = opcodes[i].low
	}
	//Set the progam counter to the location of the first opcode
	machine.PC = 0x200
}
