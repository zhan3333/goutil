package util

import (
	"bytes"
	"encoding/json"
)

// Slice 提供一系列的切片操作, 大部分操作是原地操作，返回原切片，可以用于链式操作
type Slice[T comparable] struct {
	slice []T
}

// NewSlice 新建一个集合
func NewSlice[U comparable](vs []U) *Slice[U] {
	return &Slice[U]{vs}
}

// Set 设置集合中的数据，会覆盖原有数据
func (s *Slice[T]) Set(vs []T) *Slice[T] {
	s.Reset()
	s.slice = vs
	return s
}

// Slice 返回集合中的数据
func (s *Slice[T]) Slice() []T {
	return s.slice
}

// Unique 使集合中的元素都是唯一的
func (s *Slice[T]) Unique() *Slice[T] {
	s.Set(Unique(s.Slice()))
	return s
}

// Reset 重置集合中的元素，调用后元素数为 0
func (s *Slice[T]) Reset() *Slice[T] {
	s.slice = s.slice[:0]
	return s
}

// Each 遍历集合中的元素执行传入的函数
// 会更改每一个元素的值，不会返回新的集合
func (s *Slice[T]) Each(f func(T) T) *Slice[T] {
	for k, v := range s.slice {
		s.slice[k] = f(v)
	}
	return s
}

// Filter 遍历元素，使用传入的方法进行过滤，并返回过滤后的新集合
// 传入的方法返回 false 则元素被过滤，返回 true 则会出现在结果中
func (s *Slice[T]) Filter(f func(T) bool) *Slice[T] {
	s.Set(Filter(s.Slice(), f))
	return s
}

// Reject 遍历元素，使用传入的方法进行过滤，并返回过滤后的新集合
// 与 Filter 相反，f() 返回 true 的不会出现在结果中
func (s *Slice[T]) Reject(f func(T) bool) *Slice[T] {
	s.Set(Reject(s.Slice(), f))
	return s
}

// First 返回第一个元素的指针
// 当集合为空的时候，返回 nil
func (s *Slice[T]) First() *T {
	return First(s.Slice())
}

// Last 返回集合的最后一个元素
// 不会改变集合
func (s *Slice[T]) Last() *T {
	return Last(s.Slice())
}

// Empty 集合是否为空
func (s *Slice[T]) Empty() bool {
	return Empty(s.Slice())
}

// Index 返回 i 下标对应的集合元素
// 下标不存在时，返回 nil
func (s *Slice[T]) Index(i int) *T {
	if s.Len() == 0 || i > s.Len()-1 || i < 0 {
		return nil
	}
	return &s.Slice()[i]
}

// Copy 复制集合
func (s *Slice[T]) Copy() *Slice[T] {
	dst := make([]T, s.Len())
	copy(dst, s.Slice())
	return NewSlice(dst)
}

// Merge 将传入的集合组合并到集合中
func (s *Slice[T]) Merge(ss ...*Slice[T]) *Slice[T] {
	s.Set(Merge(s.Slice(), Map(ss, func(s *Slice[T]) []T {
		return s.Slice()
	})...))
	return s
}

func (s *Slice[T]) MergeSlice(arr []T) *Slice[T] {
	s.Set(Merge(s.Slice(), arr))
	return s
}

// Reverse 反转集合
func (s *Slice[T]) Reverse() *Slice[T] {
	s.Set(Reverse(s.Slice()))
	return s
}

// Random 随机返回一个元素的指针
// 使用初始化时创建的 rand 对象
// 集合为空时，返回 nil
func (s *Slice[T]) Random() *T {
	return Random(s.Slice())
}

// Shuffle 打乱集合的顺序
// 使用初始化时设置的 rand 对象
// 使用洗牌算法，原顺序是有概率出现的
func (s *Slice[T]) Shuffle() *Slice[T] {
	s.Set(Shuffle(s.Slice()))
	return s
}

// Contains 返回集合中是否存在指定的元素
func (s *Slice[T]) Contains(v T) bool {
	return Contains(s.Slice(), v)
}

func (s *Slice[T]) ContainsAll(vs ...T) bool {
	return ContainsAll(s.Slice(), vs)
}

// ContainsCount 返回指定元素在集合中出现的次数
func (s *Slice[T]) ContainsCount(v T) int {
	return CountIf(s.Slice(), func(t T) bool {
		return t == v
	})
}

// Push 向集合的末尾添加一个元素
func (s *Slice[T]) Push(vs ...T) *Slice[T] {
	s.Set(Push(s.Slice(), vs...))
	return s
}

// Pop 弹出集合末尾的元素
// 集合为空时，返回 nil
func (s *Slice[T]) Pop() *T {
	items, last := Pop(s.Slice())
	s.Set(items)
	return last
}

// Len 返回集合的元素个数
func (s *Slice[T]) Len() int {
	return len(s.Slice())
}

func (s *Slice[T]) Equal(s2 *Slice[T]) bool {
	return Equal(s.Slice(), s2.Slice())
}

// JSON :json.Marshal 处理集合元素
func (s *Slice[T]) JSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

// JSONString 同 JSON, 但是结果会作为 string 返回
func (s *Slice[T]) JSONString() (string, error) {
	b, err := s.JSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Pretty 调试方法，返回美化的 json 字符串
func (s *Slice[T]) Pretty() string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	_ = jsonEncoder.Encode(s.Slice())
	return bf.String()
}

// Diff 返回在集合中，但是不在传入集合 c2 中的值，以集合类型返回
// 返回新的集合
func (s *Slice[T]) Diff(c2 *Slice[T]) *Slice[T] {
	return NewSlice(Diff(s.Slice(), c2.Slice()))
}

// Map 遍历集合的元素，并使用传入的方法处理元素
// 返回新的集合
// 方法不能再定义新的泛型，所以响应值只能元素类型，需要响应与输入类型不一致的可以直接用 Map() 方法
func (s *Slice[T]) Map(f func(T) T) *Slice[T] {
	return NewSlice(Map(s.Slice(), f))
}
