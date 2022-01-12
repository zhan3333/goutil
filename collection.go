package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Collection 集合，提供快速处理集合的一些方法
type Collection[T comparable] struct {
	// 集合数据
	items []T
	// 随机数对象，会被 Random(), Shuffle() 方法使用
	rand *rand.Rand
}

// Collect 新建一个集合
func Collect[T comparable](elems []T) *Collection[T] {
	rand.Seed(time.Now().Unix())
	c := Collection[T]{items: elems}
	c.rand = rand.New(rand.NewSource(time.Now().Unix()))
	return &c
}

// Set 设置集合中的数据，会覆盖原有数据
func (c *Collection[T]) Set(elems []T) {
	c.items = elems
}

// Map 遍历集合的元素，并使用传入的方法处理，返回一个新的集合
func (c *Collection[T]) Map(f func(elem T) T) *Collection[T] {
	ret := Collect([]T{})
	for _, v := range c.items {
		ret.Append(f(v))
	}
	return ret
}

// Unique 使集合中的元素都是唯一的
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

// Reset 重置集合中的元素，调用后元素数为 0
func (c *Collection[T]) Reset() {
	c.items = []T{}
}

// Filter 遍历元素，使用传入的方法进行过滤，并返回过滤后的新集合
// 传入的方法返回 false 则元素被过滤，返回 true 则会出现在结果中
func (c *Collection[T]) Filter(f func(elem T, index int) bool) *Collection[T] {
	ret := Collect([]T{})
	for k, v := range c.Items() {
		if f(v, k) {
			ret.Append(v)
		}
	}
	return ret
}

// Reject 遍历元素，使用传入的方法进行过滤，并返回过滤后的新集合
// 与 Filter 相反，f() 返回 true 的不会出现在结果中
func (c *Collection[T]) Reject(f func(elem T, index int) bool) *Collection[T] {
	ret := Collect([]T{})
	for k, v := range c.Items() {
		if !f(v, k) {
			ret.Append(v)
		}
	}
	return ret
}

// First 返回第一个元素的指针
// 当集合为空的时候，返回 nil
func (c *Collection[T]) First() *T {
	if c.Empty() {
		return nil
	}
	return &(c.Items()[0])
}

// Empty 集合是否为空
func (c *Collection[T]) Empty() bool {
	return c.Len() == 0
}

// Slice 以 slice 类型返回元素集合
func (c *Collection[T]) Slice() []T {
	return c.Items()
}

// Items 以 slice 类型返回元素集合
func (c *Collection[T]) Items() []T {
	return c.items
}

// Index 返回 i 下标对应的集合元素
// 下标不存在时，返回 nil
func (c *Collection[T]) Index(i int) *T {
	if c.Len() == 0 || i > c.Len()-1 || i < 0 {
		return nil
	}
	return &c.Items()[i]
}

// Copy 复制集合
func (c *Collection[T]) Copy() *Collection[T] {
	dst := make([]T, c.Len())
	copy(dst, c.Items())
	return Collect[T](dst)
}

// Merge 将传入的集合组合并到集合中
func (c *Collection[T]) Merge(collects ...*Collection[T]) {
	for _, collect := range collects {
		c.items = append(c.items, collect.Items()...)
	}
}

// Each 遍历集合，并使用传入的方法处理
// 当 f() 返回 error != nil 时，遍历会终止
func (c *Collection[T]) Each(f func(elem T, index int) error) error {
	for k, v := range c.items {
		if err := f(v, k); err != nil {
			return err
		}
	}
	return nil
}

// Reverse 反转集合
// 属于原地反转，时间复杂度 O(n)
func (c *Collection[T]) Reverse() {
	var i, j = 0, c.Len() - 1
	for i < j {
		c.items[i], c.items[j] = c.items[j], c.items[i]
		i++
		j--
	}
}

// Random 随机返回一个元素的指针
// 使用初始化时创建的 rand 对象
// 集合为空时，返回 nil
func (c *Collection[T]) Random() *T {
	if c.Len() == 0 {
		return nil
	}
	return c.Index(c.rand.Intn(c.Len()))
}

// Every 遍历元素调用传入的方法，如果都符合条件 f() == true，那么结果是 true
// 遍历到某个元素结果为 false 时，立即停止循环并响应结果
func (c *Collection[T]) Every(f func(elem T, index int) bool) bool {
	for k, v := range c.Items() {
		if !f(v, k) {
			return false
		}
	}
	return true
}

// Pad 使用指定的元素 val 去填充集合长度到 start
// start < 0 时在集合左侧填充
// start > 0 时在集合右侧填充
// start 为集合的最终长度，当 start <= len(集合) 时，填充不会进行
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

// Shuffle 打乱集合的顺序
// 使用初始化时设置的 rand 对象
// 使用洗牌算法，原顺序是有概率出现的
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

// Dump 调试方法，向控制台输出 json 美化后的集合
func (c *Collection[T]) Dump() {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	_ = jsonEncoder.Encode(c.Items())
	fmt.Printf("%s", bf.String())
}

// Pluck 遍历集合，使用传入的方法处理元素，并将结果作为 slice 返回
func (c *Collection[T]) Pluck(f func(elem T) any) []any {
	ret := make([]any, c.Len())
	_ = c.Each(func(e T, index int) error {
		ret[index] = f(e)
		return nil
	})
	return ret
}

// Contains 返回集合中是否存在指定的元素
func (c *Collection[T]) Contains(elem T) bool {
	for _, v := range c.Items() {
		if v == elem {
			return true
		}
	}
	return false
}

// ContainsCount 返回指定元素在集合中出现的次数
func (c *Collection[T]) ContainsCount(elem T) int {
	return CountIf(c.Items(), func(t T) bool {
		return t == elem
	})
}

// Diff 返回在集合中，但是不在传入集合 c2 中的值，以集合类型返回
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

// Append 添加元素到集合中
func (c *Collection[T]) Append(elems ...T) {
	c.items = append(c.items, elems...)
}

// Push 向集合的末尾添加一个元素
func (c *Collection[T]) Push(elems ...T) {
	c.items = append(c.items, elems...)
}

// Pop 弹出集合末尾的元素
// 集合为空时，返回 nil
func (c *Collection[T]) Pop() *T {
	if len(c.items) == 0 {
		return nil
	}
	last := c.items[len(c.items)-1]
	c.items = c.items[0 : len(c.items)-1]
	return &last
}

// Len 返回集合的元素个数
func (c *Collection[T]) Len() int {
	return len(c.items)
}

// Last 返回集合的最后一个元素
// 不会改变集合
func (c *Collection[T]) Last() *T {
	if c.Len() == 0 {
		return nil
	}
	return &(c.items[len(c.items)-1])
}

// JSON :json.Marshal 处理集合元素
func (c *Collection[T]) JSON() ([]byte, error) {
	return json.Marshal(c.Items())
}

// JSONString 同 JSON, 但是结果会作为 string 返回
func (c *Collection[T]) JSONString() (string, error) {
	b, err := c.JSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
