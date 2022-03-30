package interpreter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackPopNil(t *testing.T) {
	stack := myStack{}
	_, err := stack.pop()
	require.NotNil(t, err)
}

func TestStackPushPop(t *testing.T) {
	stack := myStack{}
	item := stackContext{}
	stack.push(&item)
	received, err := stack.pop()
	require.Nil(t, err)
	require.Equal(t, received, &item)
	require.Equal(t, stack.len(), 0)
}

func TestStackPush(t *testing.T) {
	stack := myStack{}
	item := stackContext{}
	stack.push(&item)
	require.Equal(t, stack.len(), 1)
}
