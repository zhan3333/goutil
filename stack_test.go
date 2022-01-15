package util_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	util "github.com/zhan3333/goutil"
)

func TestNewStack(t *testing.T) {
	s := util.NewStack[int]()
	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Len())
	assert.Nil(t, s.Top())
	assert.Nil(t, s.Pop())

	s.Push(1)
	assert.False(t, s.Empty())
	assert.Equal(t, 1, *s.Top())
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, 1, *s.Pop())
	assert.Equal(t, 0, s.Len())

	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, *s.Top())
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, 2, *s.Pop())
	assert.Equal(t, 1, *s.Top())
	assert.Equal(t, 1, s.Len())

	s.Clear()
	assert.Equal(t, 0, s.Len())

}
