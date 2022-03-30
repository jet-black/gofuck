package interpreter

import (
	"io"
)

type State struct {
	Mem    []byte
	Pos    int
	Input  io.Reader
	Output io.Writer
}

type Operation func(*State) error

// Config Configuration for brainfuck interpreter
type Config struct {
	Program            io.RuneReader       // brainfuck code stream. Can be string, e.g. strings.NewReader("+++")
	Input              io.Reader           // user-defined input for program. Can be stdin, e.g. bufio.NewReader(os.Stdin)
	Output             io.Writer           // program output. Can be stdout, e.g. bufio.NewWriter(os.Stdout)
	OperationsRegistry *OperationsRegistry // Registry for custom and default brainfuck operations
}
