package rb_tree

import (
	"golang.org/x/exp/constraints"
)

type RbTreeNode[T constraints.Ordered] struct {
	val     T
	isBlack bool
	left    *RbTreeNode[T]
	right   *RbTreeNode[T]
}

type RbTree[T constraints.Ordered] struct {
	root *RbTreeNode[T]
}

type Rotation int

const (
	LL Rotation = iota
	LR
	RR
	RL
)

func CreateRbTree[T constraints.Ordered]() RbTree[T] {
	return RbTree[T]{}
}

func (rbtree *RbTree[T]) Insert(val T) {
	if rbtree.root == nil {
		rbtree.root = &RbTreeNode[T]{val: val, isBlack: true}
		return
	}

	willRecolor := false
	curr := rbtree.root
	path := []*RbTreeNode[T]{curr}
	for curr != nil {
		if curr.val >= val {
			if curr.left == nil {
				curr.left = &RbTreeNode[T]{val: val}
				if !curr.isBlack {
					willRecolor = true
				}
				path = append(path, curr.left)
				break
			} else {
				curr = curr.left
				path = append(path, curr)
				continue
			}
		} else {
			if curr.right == nil {
				curr.right = &RbTreeNode[T]{val: val}
				if !curr.isBlack {
					willRecolor = true
				}
				path = append(path, curr.right)
				break
			} else {
				curr = curr.right
				path = append(path, curr)
				continue
			}
		}
	}

	if willRecolor {
		rbtree.recolor(path)
	}
}

func (rbtree *RbTree[T]) recolor(path []*RbTreeNode[T]) {

}

func (rbtree *RbTreeNode[T]) rotate(node *RbTreeNode[T]) {
	var rot Rotation
	switch rot {
	case LL:
	case LR:
	case RR:
	case RL:
	}
}

func (rbtree *RbTree[T]) Remove(val T) {

}

func (rbtree *RbTree[T]) IsValidRbTree() bool {
	if !rbtree.root.isBlack {
		return false
	}

	// stack = []RbTreeNode[T]{}
	// for curr != nil {

	// }
	return true
}
