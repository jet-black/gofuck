package main

import (
	"bytes"
	"fmt"
	"github.com/jet-black/gofuck/pkg/interpreter"
	"strings"
)

func main() {
	prog := ">>>+>>>>>+>>+>>+[<<],[\n    -[-[-[-[-[-[-[-[<+>-[>+<-[>-<-[-[-[<++[<++++++>-]<\n        [>>[-<]<[>]<-]>>[<+>-[<->[-]]]]]]]]]]]]]]]]\n    <[-<<[-]+>]<<[>>>>>>+<<<<<<-]>[>]>>>>>>>+>[\n        <+[\n            >+++++++++<-[>-<-]++>[<+++++++>-[<->-]+[+>>>>>>]]\n            <[>+<-]>[>>>>>++>[-]]+<\n        ]>[-<<<<<<]>>>>\n    ],\n]+<++>>>[[+++++>>>>>>]<+>+[[<++++++++>-]<.<<<<<]>>>>>>>>]"
	input := "program counts lines, \n words, bytes"
	inputReader := strings.NewReader(input)
	program := strings.NewReader(prog)
	var out bytes.Buffer
	ops := interpreter.NewDefaultOperationsRegistry()
	err := ops.Add('^', func(state *interpreter.State) error {
		x := state.Mem[state.Pos]
		state.Mem[state.Pos] = x * x
		return nil
	})
	if err != nil {
		return
	}
	config := &interpreter.Config{
		Program:            program,
		Input:              inputReader,
		Output:             &out,
		OperationsRegistry: ops,
	}
	svc, err := interpreter.NewInterpreter(config)
	if err != nil {
		panic(err)
	}
	err = svc.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(out.String())
	if err != nil {
		panic(err)
	}
}
