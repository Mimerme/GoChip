package main

import (
	"./chip8"
	"./debugger"
	"fmt"
	"os"
)

//These variables are only used for the debugger
var DEBUG bool
var DEBUG_PAUSE bool

//Initalize the channels to be used to cotnrol the program thread here
//We will later pass the references to the debugger & program thread from here so that the 2 threads can interact with each other
var play chan struct{}
var pause chan struct{}
var step chan struct{}

func main() {
	//Initalize some variables only used for debugging
	DEBUG = false
	DEBUG_PAUSE = false

	play = make(chan struct{})
	pause = make(chan struct{})
	step = make(chan struct{})

	//Creates a Chip8 VM
	//AKA a struct that just holds some variables
	chipVM := chip8.InitializeVM()

	//Just some greeting messages and argument handling
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

	//Creates an array of OpCodes. Each OpCode is 2 bytes.
	opcodes := chip8.ReadFile(os.Args[1])
	//Loads program into memory
	chipVM.BootstrapProgram(opcodes)

	//If the debugger is used then give the debugger UI the main thread and start the program on a seperate thread
	if DEBUG {
		go chipVM.BeginExecutionLoop(&pause, &play, &step)
		debugger.StartDebugger(chipVM, &DEBUG_PAUSE, &pause, &play, &step)
	} else {
		chipVM.BeginExecutionLoop(&pause, &play, &step)
	}
}
