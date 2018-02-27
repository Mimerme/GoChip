package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var debug_machine *Chip8

func InitDebugger(machine *Chip8) {
	debug_machine = machine
	go display_values()
}

func display_values() {
	//Clear the terminal screen for OSX and Linux only
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		fmt.Println("==PC==")
		fmt.Println(debug_machine.PC)

		//fmt.Println("==Stack==")
		//stack := debug_machine.memory[STACK_START:STACK_END]
		//for i := 0; i < len(stack); i++ {
		//	fmt.Println(stack[i])
		//}
	}
}
