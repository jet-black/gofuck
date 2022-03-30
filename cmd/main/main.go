package main

import (
	"bufio"
	"github.com/jet-black/gofuck/pkg/interpreter"
	"os"
	"strings"
)

func main() {
	prog := ">>>>+>+++>+++>>>>>+++[\n  >,+>++++[>++++<-]>[<<[-[->]]>[<]>-]<<[\n    >+>+>>+>+[<<<<]<+>>[+<]<[>]>+[[>>>]>>+[<<<<]>-]+<+>>>-[\n      <<+[>]>>+<<<+<+<--------[\n        <<-<<+[>]>+<<-<<-[\n          <<<+<-[>>]<-<-<<<-<----[\n            <<<->>>>+<-[\n              <<<+[>]>+<<+<-<-[\n                <<+<-<+[>>]<+<<<<+<-[\n                  <<-[>]>>-<<<-<-<-[\n                    <<<+<-[>>]<+<<<+<+<-[\n                      <<<<+[>]<-<<-[\n                        <<+[>]>>-<<<<-<-[\n                          >>>>>+<-<<<+<-[\n                            >>+<<-[\n                              <<-<-[>]>+<<-<-<-[\n                                <<+<+[>]<+<+<-[\n                                  >>-<-<-[\n                                    <<-[>]<+<++++[<-------->-]++<[\n                                      <<+[>]>>-<-<<<<-[\n                                        <<-<<->>>>-[\n                                          <<<<+[>]>+<<<<-[\n                                            <<+<<-[>>]<+<<<<<-[\n                                              >>>>-<<<-<-\n  ]]]]]]]]]]]]]]]]]]]]]]>[>[[[<<<<]>+>>[>>>>>]<-]<]>>>+>>>>>>>+>]<\n]<[-]<<<<<<<++<+++<+++[\n  [>]>>>>>>++++++++[<<++++>++++++>-]<-<<[-[<+>>.<-]]<<<<[\n    -[-[>+<-]>]>>>>>[.[>]]<<[<+>-]>>>[<<++[<+>--]>>-]\n    <<[->+<[<++>-]]<<<[<+>-]<<<<\n  ]>>+>>>--[<+>---]<.>>[[-]<<]<\n]"
	input := "123000"
	inputReader := strings.NewReader(input)
	program := strings.NewReader(prog)
	out := bufio.NewWriterSize(os.Stdout, 1)
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
		Output:             out,
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
	err = out.Flush()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
