package util

import "constraints"

// Map 使用传入的方法循环处理传入的 arr slice, 并将结果合并为一个新的 slice 返回
func Map[T1 any, T2 any](arr []T1, f func(T1) T2) []T2 {
	ret := make([]T2, len(arr))
	for i, elem := range arr {
		ret[i] = f(elem)
	}
	return ret
}

// Reduce 使用传入的方法聚合 arr slice
func Reduce[E any, O any](arr []E, init O, f func(O, E) O) O {
	result := init
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

// Filter 使用传入的方法过滤 arr slice
func Filter[T any](arr []T, f func(T) bool) []T {
	var ret []T
	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
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

// CountIf 返回符合条件的元素个数
func CountIf[T any](arr []T, f func(t T) bool) int {
	var count int
	for _, v := range arr {
		if f(v) {
			count++
		}
	}
	return count
}
