package chip8

import (
	"os"
)

//Stack grows down?

func display_clear() {
	//fmt.Println("This should clear the display")
	placeholder()
}

//Pop a return value off the stack and set the program counter to it
func sub_ret() {
	//fmt.Println("Return")
	machine.SP = machine.SP - 1

	//Check to see if the SP wraps
	if machine.SP == 255 {
		os.Exit(0)
	}

	return_pointer := machine.stack[machine.SP]

	machine.PC = return_pointer
}

func jump(address uint16) {
	machine.PC = address
}

func call(address uint16) {
	//fmt.Println("Call")

    machine.stack[machine.SP] = machine.PC

	machine.SP = machine.SP + 1
	machine.PC = address
}

func placeholder() {
	//fmt.Println("This is a placeholder instruction")
}

func BeginExecutionLoop() {
	machine.PC = 0x200
	//Infinite loop
	for {
		//Create a temp PC so that we can update the value after execution, but also accept PC changes from the operations
		temp_pc := machine.PC
		machine.PC += 2
		parse_opcode(machine.memory[temp_pc], machine.memory[temp_pc+1])
	}
}
