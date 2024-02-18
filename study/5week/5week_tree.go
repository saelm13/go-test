package tree

type TreeNode[T any] struct {
	Value  T
	Childs []*TreeNode[T]
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}

	t.Childs = append(t.Childs, n)
	return n
}
