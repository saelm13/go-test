package main

import "fmt"

type Node[S any] struct {
	next *Node[S]
	val  S
}

func main() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}
	root.next.next.next = &Node[int]{nil, 100}

	for n := root; n != nil; n = n.next {
		fmt.Printf("node 값: %d\n", n.val)
	}
}
