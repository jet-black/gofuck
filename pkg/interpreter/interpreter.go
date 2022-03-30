package interpreter

import (
	"errors"
	"fmt"
	"io"
)

type Interpreter struct {
	state   *State
	program *myBufferedReader
	ops     map[rune]Operation
}

const (
	LParen = '['
	RParen = ']'
)

func (interpreter *Interpreter) Execute() error {
	parenDiff := 0
	stack := myStack{
		data: make([]*stackContext, 0),
	}
	root := &stackContext{
		pos:       0,
		enterLoop: true,
		root:      true,
	}
	stack.push(root)
	for {
		tok, err := interpreter.program.ReadToken()
		if err == io.EOF {
			if parenDiff != 0 {
				return errors.New("unexpected EOF, unbalanced loop statements")
			}
			return nil
		}
		if err != nil {
			return err
		}

		err = validate(tok, interpreter.ops)
		if err != nil {
			return err
		}
		c, err := stack.get()
		if err != nil {
			return err
		}
		if tok == RParen {
			parenDiff--
			if c.root {
				return errors.New("unexpected ]")
			}
			if interpreter.state.Mem[interpreter.state.Pos] != 0 {
				interpreter.program.pos = c.pos - 1
			}
			_, err = stack.pop()
			if err != nil {
				return err
			}
		} else if tok == LParen {
			parenDiff++
			ctx := stackContext{
				pos:       interpreter.program.pos,
				enterLoop: interpreter.state.Mem[interpreter.state.Pos] != 0,
				root:      false,
			}
			stack.push(&ctx)
		} else {
			if !c.enterLoop {
				continue
			}
			op, ok := interpreter.ops[tok]
			if !ok {
				continue
			}
			err = op(interpreter.state)
			if err != nil {
				return err
			}
		}
	}
}

func validate(tok rune, ops map[rune]Operation) error {
	if tok == LParen || tok == RParen {
		return nil
	}
	if tok == '\n' || tok == '\r' || tok == '\t' || tok == ' ' {
		return nil
	}
	_, ok := ops[tok]
	if !ok {
		return fmt.Errorf("operation %c is not registered", tok)
	}
	return nil
}

func NewInterpreter(config *Config) (*Interpreter, error) {
	reader := &myBufferedReader{
		buf:        make([]rune, 0),
		underlying: config.Program,
		pos:        0,
	}
	state := &State{
		Mem:    make([]byte, 1),
		Pos:    0,
		Input:  config.Input,
		Output: config.Output,
	}
	interpreter := &Interpreter{
		state:   state,
		program: reader,
		ops:     config.OperationsRegistry.ops,
	}
	return interpreter, nil
}
