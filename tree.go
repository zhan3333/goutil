package util

type TreeNode[T any] struct {
	val   T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewTree[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{val: val}
}

func (n *TreeNode[T]) Val() T {
	return n.val
}

func (n *TreeNode[T]) Left() *TreeNode[T] {
	return n.left
}

func (n *TreeNode[T]) Right() *TreeNode[T] {
	return n.right
}

func (n *TreeNode[T]) SetLeft(val T) {
	n.left = &TreeNode[T]{val: val}
}

func (n *TreeNode[T]) SetRight(val T) {
	n.right = &TreeNode[T]{val: val}
}
