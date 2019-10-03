package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate stacker -type int -package main_test -filesuffix=_test

func Test_Push(t *testing.T) {
	var s intStack

	s.Push(1)
	assertStackEqual(t, s, []int{1})

	s.Push(2)
	assertStackEqual(t, s, []int{1, 2})
}

func Test_Pop_Peek(t *testing.T) {
	var s intStack
	s.Push(1)
	s.Push(3)
	s.Push(2)

	assertStackEqual(t, s, []int{1, 3, 2})

	v, _ := s.Pop()
	assert.Equal(t, 2, v)
	assertStackEqual(t, s, []int{1, 3})

	v, _ = s.Pop()
	assert.Equal(t, 3, v)
	assertStackEqual(t, s, []int{1})

	v, _ = s.Peek()
	assert.Equal(t, 1, v)
	assertStackEqual(t, s, []int{1})

	v, _ = s.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, s.Len())

	v, ok := s.Pop()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func Test_RotateUnrotate(t *testing.T) {
	var s intStack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assertStackEqual(t, s, []int{1, 2, 3})

	s.Rotate()
	assertStackEqual(t, s, []int{3, 1, 2})

	s.Rotate()
	assertStackEqual(t, s, []int{2, 3, 1})

	s.Unrotate()
	assertStackEqual(t, s, []int{3, 1, 2})
}

func assertStackEqual(t *testing.T, s intStack, expected []int) {
	var actual []int
	s.Walk(func(v int) {
		actual = append(actual, v)
	})
	assert.Equal(t, expected, actual)
}
