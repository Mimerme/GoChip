package main

import (
	"./chip8"
	"./debugger"
	"./io/display"
	"./io/keyboard"
	"fmt"
	"golang.org/x/image/colornames"
	"os"
)
import "github.com/faiface/pixel/pixelgl"

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
	if len(os.Args) >= 3 {
		if os.Args[2] == "d" {
			DEBUG = true
		}
	}

	//Creates an array of OpCodes. Each OpCode is 2 bytes.
	opcodes := chip8.ReadFile(os.Args[1])
	//Loads program into memory
	chipVM.BootstrapProgram(opcodes)

	//No need for channels since we're on the main thread
	var paused bool = false
	var execute_next bool = false

	//If the debugger is used then give the debugger UI the main thread and start the program on a seperate thread
	pixelgl.Run(
		func() {
			var main_window *pixelgl.Window

			if DEBUG {
				display.CreateWindow(500, 0)
				main_window = debugger.CreateWindow()
			} else {
				main_window = display.CreateWindow(0, 0)
			}

			//Initalize the pointers to the VM & Window struct
			keyboard.Initialize(main_window)

			//Start the timers on seperate threads
			//TODO: Implement pausing on these threads
			//and/or implement these on the main thread
			go chip8.StartDelayTimer(chipVM)
			go chip8.StartSoundTimer(chipVM)

			//Render loop
			for !main_window.Closed() {
				if DEBUG {
					main_window.Clear(colornames.Black)
					debugger.Render(main_window, chipVM, &paused, &execute_next)
					display.Render(main_window)
					main_window.Update()
				} else {
					main_window.Clear(colornames.Black)
					for i := 0; i < 64; i++ {
						for k := 0; k < 32; k++ {
							display.Draw(i, k)
						}
					}

					display.Render(main_window)
					main_window.Update()
				}
			}
		},
	)
	//go chipVM.BeginExecutionLoop(&pause, &play, &step)
	//debugger.StartDebugger(chipVM, &DEBUG_PAUSE, &pause, &play, &step)
	//Start the debugger paused?
	//else {
	//	chipVM.BeginExecutionLoop(&pause, &play, &step)
	//}
}
