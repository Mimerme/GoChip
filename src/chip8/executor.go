package chip8

import (
	"os"
)

//Stack grows down?

func (machine *Chip8) display_clear() {
	//fmt.Println("This should clear the display")
	machine.placeholder()
}

//Pop a return value off the stack and set the program counter to it
func (machine *Chip8) sub_ret() {
	//fmt.Println("Return")
	machine.SP = machine.SP - 1

	//Check to see if the SP wraps
	if machine.SP == 255 {
		os.Exit(0)
	}

	return_pointer := machine.stack[machine.SP]

	machine.PC = return_pointer
}

func (machine *Chip8) jump(address uint16) {
	machine.PC = address
}

func (machine *Chip8) call(address uint16) {
	//fmt.Println("Call")

	machine.stack[machine.SP] = machine.PC

	machine.SP = machine.SP + 1
	machine.PC = address
}

func (machine *Chip8) placeholder() {
	//fmt.Println("This is a placeholder instruction")
}

func (machine *Chip8) BeginExecutionLoop(pause *chan struct{}, play *chan struct{}, step *chan struct{}) {
	machine.PC = 0x200

	//Infinite loop
	for {
		machine.ExecuteStep()

		//State machine that controls the progrma thread from the debug thread using the channel references
		//Only relevant when using the debugger
		select {
		case <-(*pause):
			select {
			case <-(*play):
				break
			}
		default:
			break
		}
	}
}

func (machine *Chip8) ExecuteStep() {
	//Create a temp PC so that we can update the value after execution, but also accept PC changes from the operations
	//ie. A return instruction changes the PC, but updating the PC after it offsets the PC incorrectly by 2
	//TODO: Refine the PC handling

	temp_pc := machine.PC
	machine.PC += 2
	parse_opcode(machine.memory[temp_pc], machine.memory[temp_pc+1], machine)
}
