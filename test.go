package util

type ITest[T comparable] interface {
	Test()
	Test2(elem T) T
}

type Test[T comparable] struct {
}

func (t *Test[T]) Test2(elem T) T {
	return elem
}

func (t *Test[T]) Test() {
	print("test")
}
