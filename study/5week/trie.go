package trie

type Node struct {
	Children map[rune]*Node
	Value	string
}

func NewNode(value string) *Node{
	return &Node{
		Children: make(map[rune]*Node,
				Value: value,
		}
	

func (n *Node) GetChilds() []Node {
	rst := make([]nodeinterface.Node, len(len(n.Children)))
	idx := 0
	for i, c := range.n.Children {
		rst[idx] = C
		idx++

	}
	return idx
}


func Insert(root *Node, key string) bool {
	cur := root
	for i, c := range key {
		node := root.Children[c]
		if node == nil{
			node=NewNode(key[:i+1])
			cur.Children[c] = node
				}
			cur = node
		}
	}
	return true
}