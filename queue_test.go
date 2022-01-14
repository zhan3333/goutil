package util_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	util "github.com/zhan3333/goutil"
)

func TestNewQueue(t *testing.T) {
	q := util.NewQueue[int]()
	assert.Equal(t, 0, q.Len())
	assert.Nil(t, q.Head())
	assert.Nil(t, q.End())

	q.Push(1)
	assert.Equal(t, 1, q.Len())
	assert.Equal(t, 1, *q.Head())
	assert.Equal(t, 1, *q.End())

	q.Push(2)
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 1, *q.Head())
	assert.Equal(t, 2, *q.End())

	q.Push(3)
	q.Push(4)
	q.Push(5)

	assert.Equal(t, 1, *q.Head())

	assert.Equal(t, 1, *q.Pop())
	assert.Equal(t, 2, *q.Head())
	assert.Equal(t, 5, *q.End())
	assert.Equal(t, 4, q.Len())

	t.Log("each start")
	q.Each(func(v int) int {
		t.Log(v)
		return v * v
	})
	t.Log("each end")

	assert.Equal(t, 25, *q.End())

	q.Clear()
	assert.Equal(t, 0, q.Len())
	assert.Nil(t, q.Head())
	assert.Nil(t, q.End())
}
