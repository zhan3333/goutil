package util_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"

	"github.com/zhan3333/goutil"
)

func TestMap(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, util.Map([]int{1, 2, 3}, func(t1 int) string {
		return strconv.Itoa(t1)
	}))

	assert.Equal(t, []int{1, 4, 9}, util.Map([]int{1, 2, 3}, func(t1 int) int {
		return t1 * t1
	}))
}

func TestReduce(t *testing.T) {
	assert.Equal(t, 6, util.Reduce([]int{1, 2, 3}, 0, func(o int, e int) int {
		return o + e
	}))

	assert.Equal(t, "abc", util.Reduce([]string{"a", "b", "c"}, "", func(o string, e string) string {
		return o + e
	}))
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, util.Filter([]int{1, 2, 3, 4, 5}, func(t int) bool {
		return t < 4
	}))

	assert.Equal(t, []string{"a"}, util.Filter([]string{"a", "b", "c"}, func(t string) bool {
		return t == "a"
	}))
}

func TestCountIf(t *testing.T) {
	assert.Equal(t, 1, util.CountIf([]int{1, 2, 3}, func(t int) bool {
		return t > 2
	}))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 6, util.Sum([]int{1, 2, 3}, func(t int) int {
		return t
	}))

	assert.Equal(t, "abc", util.Sum([]string{"a", "b", "c"}, func(t string) string {
		return t
	}))
}
