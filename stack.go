package util

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{[]T{}}
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() *T {
	if s.Empty() {
		return nil
	}
	v := s.Top()
	s.items = s.items[:s.Len()-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Top() *T {
	if s.Len() == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}
