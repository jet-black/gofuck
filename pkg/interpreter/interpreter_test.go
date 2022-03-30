package interpreter

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func getInterpreter(t *testing.T, input string, program string, ops map[rune]Operation) *Interpreter {
	var buf bytes.Buffer
	registry := NewDefaultOperationsRegistry()
	for k, v := range ops {
		err := registry.Add(k, v)
		require.Nil(t, err)
	}
	config := &Config{
		Program:            strings.NewReader(program),
		Input:              strings.NewReader(input),
		Output:             &buf,
		OperationsRegistry: registry,
	}
	svc, err := NewInterpreter(config)
	require.Nil(t, err)
	return svc
}

func TestUnbalancedLoops(t *testing.T) {
	interpreter := getInterpreter(t, "", "[", make(map[rune]Operation))
	err := interpreter.Execute()
	require.NotNil(t, err)
}

func TestUnbalancedRightParen(t *testing.T) {
	interpreter := getInterpreter(t, "", "]", make(map[rune]Operation))
	err := interpreter.Execute()
	require.NotNil(t, err)
}

func TestSpacesAllowed(t *testing.T) {
	interpreter := getInterpreter(t, "", "+ +", make(map[rune]Operation))
	err := interpreter.Execute()
	require.Nil(t, err)
}

func TestDefaultOps(t *testing.T) {
	interpreter := getInterpreter(t, "", "+ +", make(map[rune]Operation))
	err := interpreter.Execute()
	require.Nil(t, err)
	require.Equal(t, interpreter.state.Mem[0], uint8(2))
}

func TestCustomOps(t *testing.T) {
	ops := map[rune]Operation{
		'!': func(state *State) error {
			state.Mem[state.Pos] += 2
			return nil
		},
	}
	interpreter := getInterpreter(t, "", "!!", ops)
	err := interpreter.Execute()
	require.Nil(t, err)
	require.Equal(t, interpreter.state.Mem[0], uint8(4))
}

func TestOpError(t *testing.T) {
	ops := map[rune]Operation{
		'!': func(state *State) error {
			return errors.New("error op")
		},
	}
	interpreter := getInterpreter(t, "", "!!", ops)
	err := interpreter.Execute()
	require.NotNil(t, err)
}

func TestUnknownOp(t *testing.T) {
	interpreter := getInterpreter(t, "", "!", make(map[rune]Operation))
	err := interpreter.Execute()
	require.NotNil(t, err)
}

func TestEndToEnd(t *testing.T) {
	prog := ">>>+>>>>>+>>+>>+[<<],[\n    -[-[-[-[-[-[-[-[<+>-[>+<-[>-<-[-[-[<++[<++++++>-]<\n        [>>[-<]<[>]<-]>>[<+>-[<->[-]]]]]]]]]]]]]]]]\n    <[-<<[-]+>]<<[>>>>>>+<<<<<<-]>[>]>>>>>>>+>[\n        <+[\n            >+++++++++<-[>-<-]++>[<+++++++>-[<->-]+[+>>>>>>]]\n            <[>+<-]>[>>>>>++>[-]]+<\n        ]>[-<<<<<<]>>>>\n    ],\n]+<++>>>[[+++++>>>>>>]<+>+[[<++++++++>-]<.<<<<<]>>>>>>>>]"
	input := "program counts lines, \n words, bytes"
	interpreter := getInterpreter(t, input, prog, make(map[rune]Operation))
	var buf bytes.Buffer
	interpreter.state.Output = &buf
	err := interpreter.Execute()
	require.Nil(t, err)
	require.Equal(t, "\t1\t5\t36\n", buf.String())
}
