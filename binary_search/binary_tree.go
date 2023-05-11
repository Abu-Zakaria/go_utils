package binary_tree

import "fmt"

type BNode struct {
	value int64
	left  *BNode
	right *BNode
}

type BinaryTree struct {
	root *BNode
}

func (b *BinaryTree) Add(number int64) {
	fmt.Println(b)
}
