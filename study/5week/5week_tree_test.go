package tree

import "testing"

func TestTreeAdd(T *testing.T) {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	root.Add("D")
	d := root.Add("d")

	b.Add("E")
	b.Add("F")

	d.Add("G")

}
