package util_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/zhan3333/goutil"
)

func TestCollect(t *testing.T) {
	c := util.Collect([]int{})
	assert.NotNil(t, c)
	assert.Equal(t, 0, c.Len())
	c2 := util.Collect([]int{1, 2, 3})
	assert.NotNil(t, c2)
	assert.Equal(t, 3, c2.Len())
}

func TestCollection_Append(t *testing.T) {
	c := util.Collect([]int{})
	assert.Equal(t, 0, c.Len())
	assert.Equal(t, []int{}, c.Items())
	c.Append(1)
	assert.Equal(t, 1, c.Len())
	assert.Equal(t, []int{1}, c.Items())
}

func TestCollection_Pop(t *testing.T) {
	c := util.Collect([]int{})
	assert.Equal(t, 0, c.Len())
	assert.Nil(t, c.Pop())
	pushElem := 1
	c.Push(pushElem)
	assert.Equal(t, 1, c.Len())
	assert.NotNil(t, c.Last())
	assert.Equal(t, pushElem, *c.Last())
}

func TestCollection_Contains(t *testing.T) {
	c := util.Collect([]int{})
	assert.False(t, c.Contains(1))
	c.Append(1)
	assert.True(t, c.Contains(1))
	assert.False(t, c.Contains(2))
}

func TestCollection_ContainsCount(t *testing.T) {
	c := util.Collect([]int{})
	assert.Equal(t, 0, c.ContainsCount(1))
	c.Append(1)
	assert.Equal(t, 1, c.ContainsCount(1))
	c.Append(1)
	assert.Equal(t, 2, c.ContainsCount(1))
}

func TestCollection_Copy(t *testing.T) {
	c := util.Collect([]int{})
	copyC := c.Copy()
	assert.Equal(t, []int{}, copyC.Items())
	c.Append(1)
	assert.Equal(t, []int{}, copyC.Items())
}

func TestCollection_Diff(t *testing.T) {
	c1 := util.Collect([]int{})
	c2 := util.Collect([]int{})
	assert.Equal(t, []int{}, c1.Diff(c2).Items())

	c1.Set([]int{1, 2, 3})
	c2.Set([]int{1})
	assert.Equal(t, []int{2, 3}, c1.Diff(c2).Items())

	c1.Set([]int{})
	c2.Set([]int{1})
	assert.Equal(t, []int{}, c1.Diff(c2).Items())

}

func TestCollection_Dump(t *testing.T) {
	c := util.Collect([]string{"a", "b", "c"})
	// print:
	// [
	//	"a",
	//	"b",
	//	"b"
	//]
	c.Dump()
}

func TestCollection_Each(t *testing.T) {
	c := util.Collect([]int{1, 2, 3})
	var calls []int
	_ = c.Each(func(elem int, index int) error {
		calls = append(calls, elem)
		return nil
	})
	assert.Equal(t, []int{1, 2, 3}, calls)

	// 测试提前退出
	calls = []int{}
	var exitErr = errors.New("exited")
	err := c.Each(func(elem int, index int) error {
		if index == 2 {
			return exitErr
		}
		calls = append(calls, elem)
		return nil
	})
	assert.ErrorIs(t, err, exitErr)
	assert.Equal(t, []int{1, 2}, calls)
}

func TestCollection_Empty(t *testing.T) {
	assert.True(t, util.Collect([]int{}).Empty())
	assert.False(t, util.Collect([]int{1}).Empty())
}

func TestCollection_Every(t *testing.T) {
	assert.True(t, util.Collect([]bool{}).Every(func(elem bool, index int) bool {
		return elem
	}))
	assert.True(t, util.Collect([]bool{true, true, true}).Every(func(elem bool, index int) bool {
		return elem
	}))
	assert.False(t, util.Collect([]bool{false, true, true}).Every(func(elem bool, index int) bool {
		return elem
	}))
}

func TestCollection_Filter(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, util.Collect([]int{1, 2, 3, 4, 5}).Filter(func(elem int, index int) bool {
		return elem <= 3
	}).Items())
	assert.Equal(t, []int{}, util.Collect([]int{1, 2, 3, 4, 5}).Filter(func(elem int, index int) bool {
		return elem < 1
	}).Items())
}

func TestCollection_First(t *testing.T) {
	assert.Nil(t, util.Collect([]int{}).First())
	assert.Equal(t, 1, *util.Collect([]int{1}).First())
}

func TestCollection_Index(t *testing.T) {
	assert.Nil(t, util.Collect([]int{}).Index(0))
	assert.Equal(t, 1, *util.Collect([]int{1}).Index(0))
}

func TestCollection_Items(t *testing.T) {
	assert.Equal(t, []int{}, util.Collect([]int{}).Items())
	assert.Equal(t, []int{1}, util.Collect([]int{1}).Items())
}

func TestCollection_JSON(t *testing.T) {
	json, err := util.Collect([]int{1}).JSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`[1]`), json)

	json, err = util.Collect([]int{}).JSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`[]`), json)
}

func TestCollection_JSONString(t *testing.T) {
	json, err := util.Collect([]int{1}).JSONString()
	assert.NoError(t, err)
	assert.Equal(t, `[1]`, json)

	json, err = util.Collect([]int{}).JSONString()
	assert.NoError(t, err)
	assert.Equal(t, `[]`, json)
}

func TestCollection_Last(t *testing.T) {
	assert.Nil(t, util.Collect([]int{}).Last())
	assert.Equal(t, 1, *util.Collect([]int{1}).Last())
	assert.Equal(t, 2, *util.Collect([]int{1, 2}).Last())
}

func TestCollection_Len(t *testing.T) {
	assert.Equal(t, 0, util.Collect([]int{}).Len())
	assert.Equal(t, 1, util.Collect([]int{1}).Len())
}

func TestCollection_Map(t *testing.T) {
	assert.Equal(t, []int{1, 4, 9}, util.Collect([]int{1, 2, 3}).Map(func(elem int) int {
		return elem * elem
	}).Items())
}

func TestCollection_Merge(t *testing.T) {
	c := util.Collect([]int{1, 2, 3})
	c.Merge(util.Collect([]int{4, 5, 6}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, c.Items())
}

func TestCollection_Pad(t *testing.T) {
	c := util.Collect([]int{})
	c.Pad(3, 0)
	assert.Equal(t, []int{0, 0, 0}, c.Items())
	c.Pad(5, 1)
	assert.Equal(t, []int{0, 0, 0, 1, 1}, c.Items())
}

func TestCollection_Pluck(t *testing.T) {
	type s struct {
		name string
	}
	c := util.Collect([]s{
		{
			name: "a",
		},
		{
			name: "b",
		},
		{
			name: "c",
		},
	})
	assert.Equal(t, []interface{}{"a", "b", "c"}, c.Pluck(func(elem s) any {
		return elem.name
	}))
}

func TestCollection_Random(t *testing.T) {
	v := util.Collect([]int{1, 2, 3}).Random()
	assert.NotNil(t, v)
	t.Log(*v)

	assert.Nil(t, util.Collect([]int{}).Random())
}

func TestCollection_Reject(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, util.Collect([]int{1, 2, 3, 4, 5}).Reject(func(elem int, index int) bool {
		return elem > 3
	}).Items())
}

func TestCollection_Reset(t *testing.T) {
	c := util.Collect([]int{})
	c.Reset()
	assert.Equal(t, []int{}, c.Items())

	c2 := util.Collect([]int{1})
	c2.Reset()
	assert.Equal(t, []int{}, c2.Items())
}

func TestCollection_Reverse(t *testing.T) {
	c := util.Collect([]int{1, 2, 3})
	c.Reverse()
	assert.Equal(t, []int{3, 2, 1}, c.Items())

	c2 := util.Collect([]int{})
	c2.Reverse()
	assert.Equal(t, []int{}, c2.Items())

	c3 := util.Collect([]int{1})
	c3.Reverse()
	assert.Equal(t, []int{1}, c3.Items())

	c4 := util.Collect([]int{1, 2, 3, 4})
	c4.Reverse()
	assert.Equal(t, []int{4, 3, 2, 1}, c4.Items())

	c5 := util.Collect([]int{1, 2})
	c5.Reverse()
	assert.Equal(t, []int{2, 1}, c5.Items())
}

func TestCollection_Shuffle(t *testing.T) {
	c := util.Collect([]int{1, 2, 3})
	c.Shuffle()
	t.Log(c.Items())
}

func TestCollection_Slice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, util.Collect([]int{1, 2, 3}).Slice())
}

func TestCollection_Unique(t *testing.T) {
	c := util.Collect([]int{1, 2, 1})
	c.Unique()
	assert.Equal(t, []int{1, 2}, c.Items())
}
