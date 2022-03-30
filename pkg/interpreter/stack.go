package interpreter

import "errors"

type myStack struct {
	data []*stackContext
}

type stackContext struct {
	pos       int
	enterLoop bool
	root      bool
}

func (s *myStack) pop() (*stackContext, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

func (s *myStack) push(ctx *stackContext) {
	s.data = append(s.data, ctx)
}

func (s *myStack) get() (*stackContext, error) {
	if s.len() == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *myStack) len() int {
	return len(s.data)
}
