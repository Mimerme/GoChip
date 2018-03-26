package chip8

import (
	"../io/display"
	"../io/keyboard"
	"fmt"
	"math/rand"
	"os"
)

var masks = [8]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}

func (machine *Chip8) generate_random(KK, x byte) {
	val := (byte)(rand.Intn(255))
	machine.GPR[x] = val & KK
}

func (machine *Chip8) draw_sprite(length, pos_x, pos_y byte) {
	snap := machine.Memory[machine.I:(machine.I + (uint16)(length))]
	for y_index, elem := range snap {
		//Generate masks each with an offset of 1 more than the previous until 0b10000000
		for index, mask := range masks {
			//If a bit exists in a memory byte draw it
			if (elem & mask) == mask {
				//Returns true if it collides with another pixel
				display.Draw((int)(pos_x)+7-index, (int)(pos_y)+(y_index))
			}
		}
	}
}

//Stack grows down?

func (machine *Chip8) display_clear() {
	//fmt.Println("This should clear the display")
	display.ClearScreen()
}

//Pop a return value off the stack and set the program counter to it
func (machine *Chip8) sub_ret() {
	machine.SP = machine.SP - 1

	//Check to see if the SP wraps
	if machine.SP == 255 {
		os.Exit(0)
	}

	return_pointer := machine.Stack[machine.SP]

	machine.PC = return_pointer
}

func (machine *Chip8) jump(address uint16) {
	machine.PC = address
}

func (machine *Chip8) call(address uint16) {
	machine.Stack[machine.SP] = machine.PC

	machine.SP = machine.SP + 1
	machine.PC = address
}

func (machine *Chip8) skip_if_equal(register uint8, value uint8) {
	if machine.GPR[register] == value {
		machine.PC += 4
	}
}

func (machine *Chip8) skip_if_not_equal(register uint8, value uint8) {
	if machine.GPR[register] != value {
		machine.PC += 4
	}
}

func (machine *Chip8) skip_if_equal_reg(r1, r2 byte) {
	if machine.GPR[r1] == machine.GPR[r2] {
		machine.PC += 4
	}
}

func (machine *Chip8) skip_if_not_equal_reg(r1, r2 byte) {
	if machine.GPR[r1] != machine.GPR[r2] {
		machine.PC += 4
	}
}

func (machine *Chip8) placeholder() {
	fmt.Println("This is a placeholder instruction")
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
	//TODO: Plz help idk what the cpu instruction cycle is

	//Poll from the keyboard
	keyboard.Check_Keys(&(machine.Keys))

	parse_opcode(machine.Memory[machine.PC], machine.Memory[machine.PC+1], machine)
}
