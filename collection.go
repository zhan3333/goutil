package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Collection[T comparable] struct {
	items []T
	rand  *rand.Rand
}

func Collect[T comparable](elems []T) *Collection[T] {
	rand.Seed(time.Now().Unix())
	c := Collection[T]{items: elems}
	c.rand = rand.New(rand.NewSource(time.Now().Unix()))
	return &c
}

func (c *Collection[T]) Set(elems []T) {
	c.items = elems
}

func (c *Collection[T]) Map(f func(elem T) T) *Collection[T] {
	ret := Collect([]T{})
	for _, v := range c.items {
		ret.Append(f(v))
	}
	return ret
}

func (c *Collection[T]) Unique() {
	ret := map[T]struct{}{}
	for _, v := range c.items {
		ret[v] = struct{}{}
	}
	c.Reset()
	for k := range ret {
		c.Append(k)
	}
}

func (c *Collection[T]) Reset() {
	c.items = []T{}
}

func (c *Collection[T]) Filter(f func(elem T, index int) bool) *Collection[T] {
	ret := Collect([]T{})
	for k, v := range c.Items() {
		if f(v, k) {
			ret.Append(v)
		}
	}
	return ret
}

func (c *Collection[T]) Reject(f func(elem T, index int) bool) *Collection[T] {
	ret := Collect([]T{})
	for k, v := range c.Items() {
		if !f(v, k) {
			ret.Append(v)
		}
	}
	return ret
}

func (c *Collection[T]) First() *T {
	if c.Empty() {
		return nil
	}
	return &(c.Items()[0])
}

func (c *Collection[T]) Empty() bool {
	return c.Len() == 0
}

func (c *Collection[T]) Slice() []T {
	return c.Items()
}

func (c *Collection[T]) Items() []T {
	return c.items
}

func (c *Collection[T]) Index(i int) *T {
	if c.Len() == 0 || i > c.Len()-1 || i < 0 {
		return nil
	}
	return &c.Items()[i]
}

func (c *Collection[T]) Copy() *Collection[T] {
	dst := make([]T, c.Len())
	copy(dst, c.Items())
	return Collect[T](dst)
}

func (c *Collection[T]) Merge(collects ...*Collection[T]) {
	for _, collect := range collects {
		c.items = append(c.items, collect.Items()...)
	}
}

func (c *Collection[T]) Each(f func(elem T, index int) error) error {
	for k, v := range c.items {
		if err := f(v, k); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection[T]) Reverse() {
	var i, j = 0, c.Len() - 1
	for i < j {
		c.items[i], c.items[j] = c.items[j], c.items[i]
		i++
		j--
	}
}

func (c *Collection[T]) Random() *T {
	if c.Len() == 0 {
		return nil
	}
	return c.Index(c.rand.Intn(c.Len()))
}

func (c *Collection[T]) Every(f func(elem T, index int) bool) bool {
	for k, v := range c.Items() {
		if !f(v, k) {
			return false
		}
	}
	return true
}

func (c *Collection[T]) Pad(start int, val T) {
	var negative bool
	if start < 0 {
		negative = true
		start = -start
	}
	if c.Len() >= start {
		return
	}
	pods := make([]T, start-c.Len())
	for i := range pods {
		pods[i] = val
	}
	if negative {
		// 负数填充左边
		c.items = append(pods, c.items...)
	} else {
		c.items = append(c.items, pods...)
	}
}

func (c *Collection[T]) Shuffle() {
	// 洗牌算法
	if c.Len() == 0 {
		return
	}
	lastI := c.Len() - 1
	for lastI > 0 {
		randI := c.rand.Intn(lastI + 1)
		c.items[randI], c.items[lastI] = c.items[lastI], c.items[randI]
		lastI--
	}
}

func (c *Collection[T]) Dump() {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	_ = jsonEncoder.Encode(c.Items())
	fmt.Printf("%s", bf.String())
}

func (c *Collection[T]) Pluck(f func(elem T) any) []any {
	ret := make([]any, c.Len())
	_ = c.Each(func(e T, index int) error {
		ret[index] = f(e)
		return nil
	})
	return ret
}

func (c *Collection[T]) Contains(elem T) bool {
	for _, v := range c.Items() {
		if v == elem {
			return true
		}
	}
	return false
}

func (c *Collection[T]) ContainsCount(elem T) int {
	return CountIf(c.Items(), func(t T) bool {
		return t == elem
	})
}

// Diff 返回在集合中，但是不在 c2 集合中的值
func (c *Collection[T]) Diff(c2 *Collection[T]) *Collection[T] {
	ret := Collect([]T{})
	m := map[T]struct{}{}
	for _, v := range c2.Items() {
		m[v] = struct{}{}
	}
	for _, v := range c.Items() {
		if _, ok := m[v]; !ok {
			ret.Append(v)
		}
	}
	return ret
}

// Append 添加元素
func (c *Collection[T]) Append(elems ...T) {
	c.items = append(c.items, elems...)
}

func (c *Collection[T]) Push(elem T) {
	c.items = append(c.items, elem)
}

func (c *Collection[T]) Pop() *T {
	if len(c.items) == 0 {
		return nil
	}
	last := c.items[len(c.items)-1]
	c.items = c.items[0 : len(c.items)-1]
	return &last
}

func (c *Collection[T]) Len() int {
	return len(c.items)
}

func (c *Collection[T]) Last() *T {
	if c.Len() == 0 {
		return nil
	}
	return &(c.items[len(c.items)-1])
}

func (c *Collection[T]) JSON() ([]byte, error) {
	return json.Marshal(c.Items())
}

func (c *Collection[T]) JSONString() (string, error) {
	b, err := c.JSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
