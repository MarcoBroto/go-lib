package avltree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Rotation uint8

const (
	LEFT Rotation = iota
	RIGHT
	LEFT_RIGHT
	RIGHT_LEFT
)

type AvlNode[T constraints.Ordered] struct {
	val     T
	left    *AvlNode[T]
	right   *AvlNode[T]
	balance int8
}

type AvlTree[T constraints.Ordered] struct {
	root *AvlNode[T]
}

func (node *AvlNode[T]) rotate(rotation Rotation) *AvlNode[T] {
	switch rotation {
	case LEFT:
		newRoot := node.right
		node.right = newRoot.left
		newRoot.left = node
		node = newRoot
	case RIGHT:
		newRoot := node.left
		node.left = newRoot.right
		newRoot.right = node
		node = newRoot
	case LEFT_RIGHT:
		newRoot := node.left.right
		node.left.right = newRoot.left
		newRoot.left = node.left
		node.left = newRoot.right
		newRoot.right = node
		node = newRoot
	case RIGHT_LEFT:
		newRoot := node.right.left
		node.right.left = newRoot.right
		newRoot.right = node.right
		node.right = newRoot.left
		newRoot.left = node
		node = newRoot
	default:
		panic("Invalid rotatoin")
	}
	return node
}

func (node *AvlNode[T]) calcBalance() int8 {
	leftH, rightH := 0, 0
	temp := node
	for temp != nil {
		leftH++
	}
	temp = node
	for temp != nil {
		rightH++
	}
	return int8(leftH - rightH)
}

func CreateAvlTree[T constraints.Ordered]() AvlTree[T] {
	return AvlTree[T]{}
}

func (tree *AvlTree[T]) Insert(val T) {
	didInsert, stack := tree.insert(val)
	if !didInsert {
		return
	}
	tree.balance(stack)
}

func (tree *AvlTree[T]) insert(val T) (bool, *[]*AvlNode[T]) {
	stack := make([]*AvlNode[T], 0, 10)
	curr := tree.root
	for curr != nil {
		stack = append(stack, curr)
		if curr.val < val {
			curr = curr.left
		} else if curr.val > val {
			curr = curr.right
		} else {
			return false, &stack
		}
	}
	return true, &stack
}

func (tree *AvlTree[T]) balance(stack *[]*AvlNode[T]) {
	for _, node := range *stack {
		fmt.Println("%?", node)
		balance := 77
		if balance > 1 {
			node.rotate(RIGHT)
		} else if balance < -1 {
			node.rotate(LEFT)
		} else {
			continue
		}
	}
}

func (tree *AvlTree[T]) Remove(val T) T {
	panic("")
}

func (tree *AvlTree[T]) Pop() T {
	panic("")
}
