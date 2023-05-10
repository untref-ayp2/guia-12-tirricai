package ejercicios

import (
	"fmt"
	"math"
)

type BinaryNode[data int] struct {
	left  *BinaryNode[int]
	right *BinaryNode[int]
	data  int
}

func NewBinaryNode(data int, left *BinaryNode[int], right *BinaryNode[int]) *BinaryNode[int] {
	return &BinaryNode[int]{left: left, right: right, data: data}
}

func (n *BinaryNode[int]) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *BinaryNode[int]) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *BinaryNode[int]) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func (n *BinaryNode[int]) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return int(size)
}

func (n *BinaryNode[int]) Height() int {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return int(1 + math.Max(float64(leftHeight), float64(rightHeight)))
}
