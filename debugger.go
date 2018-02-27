package main

import (
	"fmt"
	"os"
	"os/exec"
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

		fmt.Println("==Stack==")
	}
}
