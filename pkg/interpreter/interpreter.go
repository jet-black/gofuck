package interpreter

import (
	"errors"
	"fmt"
	"github.com/jet-black/gofuck/internal/util"
	"io"
)

type Interpreter struct {
	state   *State
	program *util.MyBufferedReader
	ops     map[rune]Operation
}

const (
	LParen = '['
	RParen = ']'
)

func (interpreter *Interpreter) loop(parenDiff int, eval bool) error {
	for {
		err := interpreter.consumeInput(eval)
		if err == io.EOF {
			if parenDiff != 0 {
				return errors.New("unexpected EOF, unbalanced loop statements")
			}
			return nil
		}
		if err != nil {
			return err
		}
		tok, err := interpreter.program.ReadToken()
		if err != nil {
			return err
		}
		if tok == LParen {
			pos := interpreter.program.Pos
			err := interpreter.loop(parenDiff+1, interpreter.state.Mem[interpreter.state.Pos] != 0)
			if err != nil {
				return err
			}
			if interpreter.state.Mem[interpreter.state.Pos] != 0 {
				interpreter.program.Pos = pos - 1
			}
		}
		if tok == RParen {
			if parenDiff == 0 {
				return errors.New("unbalanced ]")
			}
			return nil
		}
	}

}

func (interpreter *Interpreter) consumeInput(eval bool) error {
	for {
		tok, err := interpreter.program.ReadToken()
		if err != nil {
			return err
		}
		err = validate(tok, interpreter.ops)
		if err != nil {
			return err
		}
		if tok == LParen || tok == RParen {
			interpreter.program.Pos -= 1
			break
		}
		if !eval {
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
	return nil
}

func (interpreter *Interpreter) Execute() error {
	return interpreter.loop(0, true)
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
	reader := &util.MyBufferedReader{
		Buf:        make([]rune, 0),
		Underlying: config.Program,
		Pos:        0,
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
