package chip8

func InitializeVM() *Chip8 {
	machine = Chip8{}
	return &machine
}
