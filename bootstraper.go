package main

func bootstrap_program(opcodes []OpCode, machine *Chip8) {
	for i := 0; i < len(opcodes); i++ {
		//Load the code into memory
		//Each instruction is 2 bytes
		machine.memory[0x200+(i*2)] = opcodes[i].high
		machine.memory[0x200+(i*2)+1] = opcodes[i].low
	}
	//Set the progam counter to the location of the first opcode
	machine.PC = 0x200
}
