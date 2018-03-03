package chip8

import (
	"fmt"
	"os"
)

//Stack grows down?

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
	fmt.Println("Return")
	machine.SP = machine.SP - 1

	//Check to see if the SP wraps
	if machine.SP == 255 {
		os.Exit(0)
	}

	stack_offset := ((uint16)(machine.SP))
	stack_end_pointer := (uint16)(machine.memory[STACK_START+stack_offset])
	stack_end_pointer_2 := (uint16)(machine.memory[STACK_START+stack_offset-1])

	machine.PC = (stack_end_pointer_2 << 4) | stack_end_pointer
}

func jump(address uint16) {
	machine.PC = address
}

func call(address uint16) {
	fmt.Println("Call")
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
		//Create a temp PC so that we can update the value after execution, but also accept PC changes from the operations
		temp_pc := machine.PC
		machine.PC += 2
		parse_opcode(machine.memory[temp_pc], machine.memory[temp_pc+1])
	}
}
