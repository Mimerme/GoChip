package main

import "fmt"
import "os"

func main() {
	chipVM := InitializeVM()

	fmt.Println("Chip 8 emu")
	if len(os.Args) < 2 {
		fmt.Println("Plz specify program file")
		return
	}

	if len(os.Args) == 3 {
		if os.Args[3] == "d" {
			fmt.Println("Running debugger")
			InitDebugger(chipVM)
		}
	}

	//Split the file into segments of 2 bytes
	opcodes := read_file(os.Args[1])
	//Load program into memory
	bootstrap_program(opcodes, chipVM)

	begin_execution_loop()
}
