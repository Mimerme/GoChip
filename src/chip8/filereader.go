package chip8

import "os"

type OpCode struct {
	high byte
	low  byte
}

//Opcodes are 2 bytes long

func ReadFile(filepath string) []OpCode {
	f, err := os.Open(filepath)
	file_info, err := f.Stat()
	total_opcodes := file_info.Size() / 2

	opcode_list := make([]OpCode, total_opcodes)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	var i int64
	for i = 0; i < total_opcodes; i++ {
		//Since go can only read files one byte at a time (and opcodes are 2 bytes)
		//Load in two bytes and OR them with each other
		opcode_buffer := make([]byte, 2)

		_, err = f.Read(opcode_buffer)

		var opcode OpCode
		opcode = OpCode{opcode_buffer[0], opcode_buffer[1]}
		opcode_list[i] = opcode
	}
	return opcode_list

}
