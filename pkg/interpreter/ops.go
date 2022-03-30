package interpreter

import (
	"errors"
	"fmt"
	"io"
)

func Plus(state *State) error {
	state.Mem[state.Pos]++
	return nil
}

func Minus(state *State) error {
	state.Mem[state.Pos]--
	return nil
}

func Read(state *State) error {
	b := make([]byte, 1)
	_, err := state.Input.Read(b)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	state.Mem[state.Pos] = b[0]
	return nil
}

func Print(state *State) error {
	b := make([]byte, 1)
	b[0] = state.Mem[state.Pos]
	_, err := state.Output.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func Right(state *State) error {
	state.Pos += 1
	if state.Pos >= len(state.Mem) {
		state.Mem = append(state.Mem, 0)
	}
	return nil
}

func Left(state *State) error {
	state.Pos -= 1
	if state.Pos < 0 {
		return errors.New("out of bounds")
	}
	return nil
}

type OperationsRegistry struct {
	ops map[rune]Operation
}

func (reg *OperationsRegistry) Add(opName rune, operation Operation) error {
	_, ok := reg.ops[opName]
	if ok {
		return fmt.Errorf("operation %c is already registered", opName)
	}
	reg.ops[opName] = operation
	return nil
}

func (reg *OperationsRegistry) Remove(opName rune) {
	delete(reg.ops, opName)
}

func NewDefaultOperationsRegistry() *OperationsRegistry {
	ops := make(map[rune]Operation)
	ops['+'] = Plus
	ops['-'] = Minus
	ops['>'] = Right
	ops['<'] = Left
	ops['.'] = Print
	ops[','] = Read
	return &OperationsRegistry{
		ops: ops,
	}
}
