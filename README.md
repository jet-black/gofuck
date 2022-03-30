Gof\*ck - a streaming branf\*ck interpreter
================================

## Installation

Add line to your `go.mod` file:

```
require github.com/jet-black/gofuck v0.0.3
```

## Quick start

Demo code that counts number of lines, words and bytes for a given input.

```go
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

```

## What is brainf\*ck?

https://en.wikipedia.org/wiki/Brainfuck

## Where can I find brainf\*ck code examples?

http://brainfuck.org/

## Implementation details

### Streaming manner
Gof\*ck provides brainf\*ck interpreter with streaming support.
It means your program can be generated and processed on-the fly.
To execute your streaming code, you should provide `io.RuneReader`

### Support for custom operations
You can write your own operations, or even delete default operations, if you need:
```
func PlusTwo(state *State) error {
	state.Mem[state.Pos] += 2
	return nil
}
ops := interpreter.NewDefaultOperationsRegistry()
ops.Remove('+')
err := ops.Add('^', PlusTwo)
```






