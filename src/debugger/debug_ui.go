package debugger

import "github.com/marcusolsson/tui-go"
import "fmt"
import "../chip8/"

func StartDebugger(chip8VM *(chip8.Chip8), DEBUG_PAUSE *bool, pause *chan struct{}, play *chan struct{}, step *chan struct{}) {
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
		tui.NewLabel("<Tab> to cycle through panels. <P> to pause execution. <D> Execute instruction"),
	)
	footer.SetSizePolicy(tui.Minimum, tui.Maximum)

	registers := tui.NewVBox(reg_label)
	stack := tui.NewVBox(stack_label)
	footer_box := tui.NewVBox(footer)

	status_bar := tui.NewStatusBar("Registers & Stack")
	status_bar.SetPermanentText("Executing Program...")

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

	ui.SetKeybinding("P", func() {
		*(DEBUG_PAUSE) = !(*(DEBUG_PAUSE))
		if *DEBUG_PAUSE {
			status_bar.SetPermanentText("Execution Paused")
			(*pause) <- struct{}{}
		} else {
			status_bar.SetPermanentText("Executing Program...")
			(*play) <- struct{}{}
		}
	})

	ui.SetKeybinding("D", func() {
		//Run a step
		if *DEBUG_PAUSE {
			status_bar.SetPermanentText("Stepped over instruction")
			(*chip8VM).ExecuteStep()
		}
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
