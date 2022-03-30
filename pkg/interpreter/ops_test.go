package interpreter

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func makeDefaultState(input string) *State {
	var buf bytes.Buffer
	return &State{
		Mem:    make([]byte, 1),
		Pos:    0,
		Input:  strings.NewReader(input),
		Output: &buf,
	}
}

func TestReadOp(t *testing.T) {
	input := "1"
	state := makeDefaultState(input)
	err := Read(state)
	require.Nil(t, err)
	require.Equal(t, string(state.Mem[:1]), input)
}

func TestPrintOp(t *testing.T) {
	input := "1"
	state := makeDefaultState(input)
	var buf bytes.Buffer
	state.Output = &buf
	err := Read(state)
	require.Nil(t, err)
	err = Print(state)
	require.Nil(t, err)
	require.Equal(t, buf.String(), input)
}

func TestLeftOp(t *testing.T) {
	state := makeDefaultState("")
	state.Pos = 1
	err := Left(state)
	require.Nil(t, err)
	require.Equal(t, state.Pos, 0)
}

func TestLeftOpBound(t *testing.T) {
	state := makeDefaultState("")
	state.Pos = 0
	err := Left(state)
	require.NotNil(t, err)
}

func TestRightOp(t *testing.T) {
	state := makeDefaultState("")
	state.Pos = 0
	err := Right(state)
	require.Nil(t, err)
	require.Equal(t, state.Pos, 1)
	require.Len(t, state.Mem, 2)
}

func TestMinusOp(t *testing.T) {
	state := makeDefaultState("")
	err := Minus(state)
	require.Nil(t, err)
	require.Equal(t, state.Mem[0], uint8(255))
}

func TestPlusOp(t *testing.T) {
	state := makeDefaultState("")
	err := Plus(state)
	require.Nil(t, err)
	require.Equal(t, state.Mem[0], uint8(1))
}

func TestRegistryDefaultOps(t *testing.T) {
	reg := NewDefaultOperationsRegistry()
	require.Contains(t, reg.ops, '+')
	require.Contains(t, reg.ops, '-')
	require.Contains(t, reg.ops, '>')
	require.Contains(t, reg.ops, '<')
	require.Contains(t, reg.ops, '.')
	require.Contains(t, reg.ops, ',')
}

func TestRegistryUniqueOps(t *testing.T) {
	reg := NewDefaultOperationsRegistry()
	err := reg.Add('.', func(state *State) error {
		return nil
	})
	require.NotNil(t, err)
}

func TestRegistryAddOps(t *testing.T) {
	reg := NewDefaultOperationsRegistry()
	err := reg.Add('!', func(state *State) error {
		return nil
	})
	require.Nil(t, err)
	require.Contains(t, reg.ops, '!')
}

func TestRegistryRemoveOps(t *testing.T) {
	reg := NewDefaultOperationsRegistry()
	reg.Remove('+')
	require.NotContains(t, reg.ops, '+')
}
