package util

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"reflect"
	"sort"
)

// Contains 是否包含
func Contains[T comparable](arr []T, v T) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

// ContainsAny 包含任意一个
func ContainsAny[T comparable](arr []T, vs ...T) bool {
	m := map[T]bool{}
	for _, a := range arr {
		m[a] = true
	}
	for _, v := range vs {
		if m[v] {
			return true
		}
	}
	return false
}

// ContainsAll 包含所有传入的参数
func ContainsAll[T comparable](arr []T, vs ...T) bool {
	m := map[T]bool{}
	for _, a := range arr {
		m[a] = true
	}
	for _, v := range vs {
		if !m[v] {
			return false
		}
	}
	return true
}

// Unique 去重, 保持原来顺序
func Unique[T comparable](arr []T) []T {
	var ret []T
	var m = map[T]int{}
	for _, v := range arr {
		m[v]++
	}
	for _, v := range arr {
		if m[v] > 0 {
			ret = append(ret, v)
			m[v] = 0
		}
	}
	if len(ret) == 0 {
		return []T{}
	}
	return ret
}

// Map 遍历数组，返回新的数组
func Map[T, U any](arr []T, f func(T) U) []U {
	var ret []U
	for _, v := range arr {
		ret = append(ret, f(v))
	}
	if len(ret) == 0 {
		return []U{}
	}
	return ret
}

// Reduce 遍历数组，返回一个值
func Reduce[I, R any](arr []I, f func(R, I) R) R {
	var ret R
	for _, v := range arr {
		ret = f(ret, v)
	}
	return ret
}

// Filter 遍历数组，按照传入的方法过滤数组，返回新的数组
func Filter[T any](arr []T, f func(T) bool) []T {
	var ret []T
	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}
	if len(ret) == 0 {
		return []T{}
	}
	return ret
}

// Reject 遍历数组，按照传入的方法过滤数组，返回新的数组
func Reject[T any](arr []T, f func(T) bool) []T {
	return Filter(arr, func(v T) bool {
		return !f(v)
	})
}

func First[T any](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
	return &arr[0]
}

func Last[T any](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
	return &arr[len(arr)-1]
}

func Empty[T any](arr []T) bool {
	return len(arr) == 0
}

func Merge[T any](arr []T, arrs ...[]T) []T {
	for _, v := range arrs {
		arr = append(arr, v...)
	}
	if len(arr) == 0 {
		return []T{}
	}
	return arr
}

func Reverse[T any](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func CountIf[T any](arr []T, f func(T) bool) int {
	var sum int
	for _, v := range arr {
		if f(v) {
			sum++
		}
	}
	return sum
}

func Random[T any](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
	return &arr[rand.Intn(len(arr))]
}

// Shuffle 打乱数组
func Shuffle[T any](arr []T) []T {
	// 洗牌算法
	if len(arr) == 0 {
		return arr
	}
	lastI := len(arr) - 1
	for lastI > 0 {
		randI := rand.Intn(lastI + 1)
		arr[randI], arr[lastI] = arr[lastI], arr[randI]
		lastI--
	}
	return arr
}

// Diff 差集, 存在 arr1 中但是不存在于 arr2 中的元素
func Diff[T comparable](arr1, arr2 []T) []T {
	var ret []T
	var m = map[T]bool{}
	for _, v := range arr2 {
		m[v] = true
	}
	for _, v := range arr1 {
		if !m[v] {
			ret = append(ret, v)
		}
	}
	if len(ret) == 0 {
		return []T{}
	}
	return ret
}

func Push[T any](arr []T, vs ...T) []T {
	arr = append(arr, vs...)
	return arr
}

func Pop[T any](arr []T) ([]T, *T) {
	if len(arr) == 0 {
		return arr, nil
	}
	last := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr, &last
}

type Sumable interface {
	constraints.Integer | constraints.Float | string
}

// Sum 求和
func Sum[T any, U Sumable](arr []T, f func(t T) U) U {
	var i U
	for _, v := range arr {
		i += f(v)
	}
	return i
}

func Equal[T any](arr1, arr2 []T) bool {
	return reflect.DeepEqual(arr1, arr2)
}

type sortData[T constraints.Ordered] struct {
	slice []T
}

func (s sortData[T]) Len() int {
	return len(s.slice)
}

func (s sortData[T]) Less(i, j int) bool {
	return s.slice[i] < s.slice[j]
}

func (s sortData[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func Sort[T constraints.Ordered](arr []T) []T {
	data := sortData[T]{arr}
	sort.Sort(data)
	return data.slice
}
