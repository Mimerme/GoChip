package main

import (
	"./chip8"
	"./chip8_debugger"
	"fmt"
	"os"
)

var DEBUG bool

func main() {
	DEBUG = false

	chipVM := chip8.InitializeVM()

	fmt.Println("Chip 8 VM Initialized")
	if len(os.Args) < 2 {
		fmt.Println("Plz specify program file")
		return
	}
	if len(os.Args) == 3 {
		if os.Args[2] == "d" {
			DEBUG = true
		}
	}

	//Split the file into segments of 2 bytes
	opcodes := chip8.ReadFile(os.Args[1])
	//Load program into memory
	chip8.BootstrapProgram(opcodes, chipVM)

	if DEBUG {
		go chip8.BeginExecutionLoop()
		chip8_debugger.StartDebugger(chipVM)
	} else {
		chip8.BeginExecutionLoop()
	}
}
