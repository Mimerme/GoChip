# GoChip
Chip 8 Emulator written in Go

## Structure
This is how the emulator is structured

```
   |Main| <-> |FileReader|
      |
      ↓                       
|Bootstrapper|---------------------                 
      |                           |
      |                           |
      ↓                           ↓
|Interpreter| -> |Executor| -> [CHIP8_VM]
      ↑_____loop_____|
```

- The `Main` function (Chip8.go) begins with asking `FileReader` to parse the binary of the program and then passes control to the `bootstraper`
- The `bootstraper` loads the program into memory (located in `CHIP8_VM`) and then begins the execution loop
- Execution loop
  - The `interpreter` determines which opcode it is reading and then passes control onto the `executor`
  - The `executor` can either edit the `CHIP8_VM` machine state or execute code on the host computer
- `CHIP8_VM` is available globally 

### Project Files
- `src` - contains the actual emulator files
   - `chip8` - core emulator package
      - `chip8.go`
         - `InitalizeVM` (Load chip8 character font from 0-F)
         - `Chip8` {struct} (Structure of registers and memory)
         - `StartDelayTimer` (Start a seperate thread for decresing the delay timer)
         - `StartSoundTimer` (Start a seperate thread for decresing the sound timer)
      - `executor.go` - Contains implementations for all of the OpCodes. As well as the following functions
         - `BeginExecutionLoop`
         - `ExecuteStep`
      - `interpreter.go` - Parses the OpCodes and executes its corresponding implementation in `executor.go`
      - `bootstraper.go` - load program opcodes into memory address 0x200
      - `filereader.go` - Creates an array of OpCode structs
   - `io` - I/O package for polling a keyboard and displaying the screen
      - `display` - display sub-package
         - `CreateWindow` (Creates an OpenGL Window)
         - `Draw` - enable a pixel at a certain location
         - `Render` - draw all of the pixels stored in memory
         - `ClearScreen` - clear all pixels in memory
      - `keyboard` - keyboard sub-package
         - `Check_keys` - Poll the keyboard and store the result into an array
   - `debugger` - debugger package
   - `debug_ui.go` - Creates the debugger UI
   - `draw_pannels.go` - Draw the subsections in the debugger UI
- `programs` - contains test/sample programs
