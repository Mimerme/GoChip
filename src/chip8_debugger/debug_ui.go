package chip8_debugger

import "github.com/marcusolsson/tui-go"
import "fmt"
import "../chip8/"

func StartDebugger(chip8VM *chip8.Chip8) {
	fmt.Println("Starting the debugger")

	reg_label := tui.NewVBox(
		tui.NewLabel("Registers"),
	)
	reg_label.SetSizePolicy(tui.Minimum, tui.Maximum)

	stack_label := tui.NewVBox(
		tui.NewLabel("Stack"),
	)
	stack_label.SetSizePolicy(tui.Minimum, tui.Maximum)

	footer := tui.NewVBox(
		tui.NewLabel("<Tab> to cycle through panels"),
	)
	footer.SetSizePolicy(tui.Minimum, tui.Maximum)

	registers := tui.NewVBox(reg_label)
	stack := tui.NewVBox(stack_label)
	footer_box := tui.NewVBox(footer)

	status_bar := tui.NewStatusBar("Registers & Stack")
	status_bar.SetPermanentText("Chip 8 Debugging Tools")

	main_panel := tui.NewVBox(
		status_bar,
		tui.NewHBox(registers, stack),
		footer_box,
	)

	registers.SetBorder(true)
	stack.SetBorder(true)

	//Draw the stack
	draw_stack(stack, nil)
	draw_registers(registers, &chip8VM.GPR, &chip8VM.I, &chip8VM.PC, &chip8VM.SP)

	ui, err := tui.New(main_panel)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
