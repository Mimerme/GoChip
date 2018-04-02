package chip8

import (
	"../io/display"
	"../io/keyboard"
	"fmt"
	"math/rand"
	"os"
)

var masks = [8]byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}

//Stack grows down?
//00E0
func (machine *Chip8) CLS() {
	//fmt.Println("This should clear the display")
	display.ClearScreen()
	machine.PC += 2
}

//00EE
//Pop a return value off the stack and set the program counter to it
func (machine *Chip8) RET() {
	machine.SP = machine.SP - 1

	//Check to see if the SP wraps
	if machine.SP == 255 {
		os.Exit(0)
	}

	return_pointer := machine.Stack[machine.SP]

	machine.PC = return_pointer
}

//1nnn
func (machine *Chip8) JP(address uint16) {
	machine.PC = address
}

//2nnn
func (machine *Chip8) CALL(address uint16) {
	machine.Stack[machine.SP] = machine.PC

	machine.SP = machine.SP + 1
	machine.PC = address
}

//3xkk
func (machine *Chip8) SE(register uint8, value uint8) {
	machine.PC += 2
	if machine.GPR[register] == value {
		machine.PC += 2
	}
}

//4xkk
func (machine *Chip8) SNE(register uint8, value uint8) {
	machine.PC += 2
	if machine.GPR[register] != value {
		machine.PC += 2
	}
}

//5xy0
func (machine *Chip8) SE_REG(r1, r2 byte) {
	machine.PC += 2
	if machine.GPR[r1] == machine.GPR[r2] {
		machine.PC += 2
	}
}

//9xy0
func (machine *Chip8) SNE_REG(r1, r2 byte) {
	machine.PC += 2
	if machine.GPR[r1] != machine.GPR[r2] {
		machine.PC += 2
	}
}

//6xkk
func (machine *Chip8) LD(nib_2, nib_3, nib_4 byte) {
	machine.GPR[nib_2] = (nib_3 << 4) | nib_4
	machine.PC += 2
}

//7xkk
func (machine *Chip8) ADD(nib_2, nib_3, nib_4 byte) {
	machine.GPR[nib_2] = machine.GPR[nib_2] + ((nib_3 << 4) | nib_4)
	machine.PC += 2
}

//8xy0
func (machine *Chip8) LD_REG(nib_2, nib_3 byte) {
	machine.GPR[nib_2] = machine.GPR[nib_3]
	machine.PC += 2
}

//8xy1
func (machine *Chip8) OR(nib_2, nib_3 byte) {
	machine.GPR[nib_2] = (machine.GPR[nib_3]) | machine.GPR[nib_2]
	machine.PC += 2
}

//8xy2
func (machine *Chip8) OR(nib_2, nib_3 byte) {
	machine.GPR[nib_2] = (machine.GPR[nib_3]) & machine.GPR[nib_2]
	machine.PC += 2
}

//8xy3
func (machine *Chip8) XOR(nib_2, nib_3 byte) {
	machine.GPR[nib_2] = (machine.GPR[nib_2]) ^ machine.GPR[nib_3]
	machine.PC += 2
}

//8xy4
func (machine *Chip8) ADD(nib_2, nib_3 byte) {
	var left int64 = machine.GPR[nib_2]
	var right int64 = machine.GPR[nib_3]
	if left+right > 0xFF {
		machine.GPR[0xF] = 1
	}

	machine.GPR[nib_2] = machine.GPR[nib_2] + machine.GPR[nib_3]
	machine.PC += 2
}

//8xy5
func (machine *Chip8) SUB(nib_2, nib_3 byte) {
	var left byte = machine.GPR[nib_2]
	var right byte = machine.GPR[nib_3]
	if left > right {
		machine.GPR[0xF] = 1
	} else {
		machine.GPR[0xF] = 0
	}

	machine.GPR[nib_2] = left - right
	machine.PC += 2
}

//8xy6
func (machine *Chip8) SHR(nib_2, nib_3 byte) {

	//Apply a mask to get the least sig bit
	if (machine.GPR[nib_2] & 0x01) == 0x01 {
		machine.GPR[0xF] = 1
	} else {
		machine.GPR[0xF] = 0
	}
	//Divide by 2
	machine.GPR[nib_2] = machine.GPR[nib_2] >> 1
	machine.PC += 2
}

//8xy5
func (machine *Chip8) SUBN(nib_2, nib_3 byte) {
	var left byte = machine.GPR[nib_2]
	var right byte = machine.GPR[nib_3]
	if right > left {
		machine.GPR[0xF] = 1
	} else {
		machine.GPR[0xF] = 0
	}

	machine.GPR[nib_2] = right - left
	machine.PC += 2
}

//8xyE
func (machine *Chip8) SHL(nib_2, nib_3 byte) {

	//Apply a mask to get the least sig bit
	if (machine.GPR[nib_2] & 0x01) == 0x01 {
		machine.GPR[0xF] = 1
	} else {
		machine.GPR[0xF] = 0
	}
	//Divide by 2
	machine.GPR[nib_2] = machine.GPR[nib_2] << 1
	machine.PC += 2
}

//8xyE
func (machine *Chip8) SHL(nib_2, nib_3 byte) {

	//Apply a mask to get the least sig bit
	if (machine.GPR[nib_2] & 0x01) == 0x01 {
		machine.GPR[0xF] = 1
	} else {
		machine.GPR[0xF] = 0
	}
	//Divide by 2
	machine.GPR[nib_2] = machine.GPR[nib_2] << 1
	machine.PC += 2
}

//9xy0
func (machine *Chip8) SNE_REG(nib_2, nib_3 byte) {
	machine.PC += 2
	if machine.GPR[nib_2] != machine.GPR[nib_3] {
		machine.PC += 2
	}
}

//Annn
func (machine *Chip8) LD_I(nib_2, nib_3 byte) {
	machine.I = ((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4)
	machine.PC += 2
}

//Bnnn
func (machine *Chip8) JP_V0(nib_2, nib_3 byte) {
	machine.PC = ((uint16)(nib_2) << 8) | ((uint16)(nib_3) << 4) | (uint16)(nib_4)
	machine.PC += (uint16)(machine.GPR[0])
}

//Cxkk
func (machine *Chip8) RND(KK, x byte) {
	val := (byte)(rand.Intn(255))
	machine.GPR[x] = val & KK
	machine.PC += 2
}

//Dxyn
func (machine *Chip8) DRW(length, pos_x, pos_y byte) {
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
	machine.PC += 2
}

//Ex9E
func (machine *Chip8) SKP(nib_2 byte) {
	machine.PC += 2
	if machine.Keys[nib_2] {
		machine.PC += 2
	}
}

//ExA1
func (machine *Chip8) SKNP(nib_2 byte) {
	machine.PC += 2
	if !machine.Keys[nib_2] {
		machine.PC += 2
	}
}

//Fx07
//Save the delay timer
func (machine *Chip8) STR_DT(nib_2 byte) {
	machine.GPR[nib_2] = machine.DT
	machine.PC += 2
}

//Fx0A
//TODO: Blocks execution until a key is pressed
func (machine *Chip8) BLOCK_KEY(nib_2 byte) {
	placeholder()
	machine.PC += 2
}

//Fx15
//load into the delay timer
func (machine *Chip8) LD_DT(nib_2 byte) {
	machine.DT = machine.GPR[nib_2]
	machine.PC += 2
}

//Fx18
//load into the sound timer
func (machine *Chip8) LD_ST(nib_2 byte) {
	machine.ST = machine.GPR[nib_2]
	machine.PC += 2
}

//Fx1E
func (machine *Chip8) ADD_I(nib_2 byte) {
	machine.I = machine.I + (uint16)(machine.GPR[nib_2])
	machine.PC += 2
}

//Fx29
//Fx33
//Fx55
//Fx65

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
