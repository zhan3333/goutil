package util

type QNode[T any] struct {
	val  T
	next *QNode[T]
}

type Queue[Q any] struct {
	head *QNode[Q]
	end  *QNode[Q]
	len  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (l *Queue[Q]) Push(v Q) {
	var node = &QNode[Q]{
		val:  v,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.end = node
	} else {
		l.end.next = node
		l.end = l.end.next
	}

	l.len++
}

func (l *Queue[Q]) Pop() *Q {
	if l.head == nil {
		return nil
	}
	v := l.head
	l.head = l.head.next
	if l.head == nil {
		l.end = nil
	}
	l.len--
	return &v.val
}

func (l *Queue[Q]) Clear() {
	l.head = nil
	l.end = nil
	l.len = 0
}

func (l *Queue[Q]) Len() int {
	return l.len
}

func (l *Queue[Q]) Head() *Q {
	if l.head == nil {
		return nil
	}
	return &l.head.val
}

func (l *Queue[Q]) End() *Q {
	if l.end == nil {
		return nil
	}
	return &l.end.val
}

func (l *Queue[Q]) Each(f func(v Q) Q) {
	p := l.head
	for p != nil {
		p.val = f(p.val)
		p = p.next
	}
}
