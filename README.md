# GoChip
Chip 8 Emulator written in Go

## Structure
This is how the project is structured

```
   |Main| <-> |FileReader|
      |
      ↓                       [CHIP8_VM]
|Bootstrapper|                 
      |                       
      |     
      ↓    
|Interpreter| -> |Executor|
      ↑_____loop_____|
```

- The `Main` loop (Chip8.go) begins with asking `FileReader` to parse the binary of the program and then passes control to the `bootstraper`
- The `bootstraper` loads the program into memory and then begins the execution loop
- Execution loop
  - The `interpreter` determines which opcode it is reading and then passes control onto the `executor`
  - The `executor` can either edit the `CHIP8_VM` machine state or execute code on the host computer
- `CHIP8_VM` is available globally 

### Project Files
- `core_emu` - contains the actual emulator files
- `programs` - contains test/sample programs
- `debug` - contains the code for the Chip 8 debugging tools
