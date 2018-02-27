package main

import "fmt"

//Stack grows down?

var machine Chip8

func InitializeVM() *Chip8 {
	machine = Chip8{}
	return &machine
}

func display_clear() {
	fmt.Println("This should clear the display")
	placeholder()
}

//Pop a return value off the stack and set the program counter to it
func sub_ret() {
	stack_offset := ((uint16)(machine.SP))
	stack_end_pointer := (uint16)(machine.memory[STACK_START+stack_offset])
	stack_end_pointer_2 := (uint16)(machine.memory[STACK_START+stack_offset-1])

	machine.PC = (stack_end_pointer_2 << 4) | stack_end_pointer
	machine.SP = machine.SP - 1
}

func jump(address uint16) {
	machine.PC = address
}

func call(address uint16) {
	machine.SP = machine.SP + 1
	upper_pc := ((machine.PC & 0xFF00) >> 8)
	lower_pc := (machine.PC & 0x00FF)

	machine.memory[STACK_START+(uint16)(machine.SP)] = (byte)(upper_pc)
	machine.memory[STACK_START+(uint16)(machine.SP+1)] = (byte)(lower_pc)
	machine.PC = address
}

func placeholder() {
	fmt.Println("This is a placeholder instruction")
}

func begin_execution_loop() {
	machine.PC = 0x200
	//Infinite loop
	for {
		parse_opcode(machine.memory[machine.PC], machine.memory[machine.PC+1])
	}
	machine.PC += 1
}
