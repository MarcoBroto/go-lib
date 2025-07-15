package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type RbTreeNode[T constraints.Ordered] struct {
	val     T
	isBlack bool
	left    *RbTreeNode[T]
	right   *RbTreeNode[T]
	parent  *RbTreeNode[T]
}

type RbTree[T constraints.Ordered] struct {
	root     *RbTreeNode[T]
	capacity uint
}

type Rotation uint8

const (
	LL Rotation = iota
	LR
	RR
	RL
)

func CreateRbTree[T constraints.Ordered]() *RbTree[T] {
	return &RbTree[T]{}
}

func printNode[T constraints.Ordered](node *RbTreeNode[T], depth uint) {
	if node == nil {
		return
	}
	printNode(node.left, depth+1)
	var colorStr string = "B"
	if !node.isBlack {
		colorStr = "R"
	}
	fmt.Printf("%d:(%d,%s) ", depth, any(node.val), colorStr)
	printNode(node.right, depth+1)
}

func (rbtree *RbTree[T]) Print() {
	fmt.Print("Tree contents: ")
	printNode(rbtree.root, 1)
	fmt.Println()
}

func (rbtree *RbTree[T]) Insert(val T) {
	// If root doesn't exist, create root node
	if rbtree.root == nil {
		rbtree.root = &RbTreeNode[T]{val: val, isBlack: true, parent: nil}
		rbtree.capacity++
		return
	}

	curr := rbtree.root
	for curr != nil {
		if curr.val > val {
			if curr.left == nil {
				// Insert new red left child node
				curr.left = &RbTreeNode[T]{val: val, isBlack: false, parent: curr}
				rbtree.capacity++
				// Save inserted node
				curr = curr.left
				break
			} else {
				// Traverse to left child node
				curr = curr.left
			}
		} else if curr.val < val {
			if curr.right == nil {
				// Insert new red right child node
				curr.right = &RbTreeNode[T]{val: val, isBlack: false, parent: curr}
				rbtree.capacity++
				// Save inserted node
				curr = curr.right
				break
			} else {
				// Traverse to right child node
				curr = curr.right
			}
		} else {
			// Skip duplicate value
			return
		}
	}

	// TODO: remove assertion - Assert newly created node is not nil
	if curr == nil {
		panic("Insertion error: Newly created node should not be nil")
	}

	// Recolor tree and fix violations resulting beginning from the newly inserted node
	rbtree.recolor(curr)
}

func (rbtree *RbTree[T]) rotate(node *RbTreeNode[T], rotation Rotation) *RbTreeNode[T] {
	if node == nil {
		// Assert provided node is not nil
		panic("rotate error: provided node cannot be nil")
	} else if node.isBlack {
		// Assert current node is not black
		panic("rotate error: Provided node is black. 'rotate' should only be called when the provided node is a newly inserted red node")
	}

	var parent *RbTreeNode[T] = node.parent
	if parent == nil {
		// Assert parent node is not nil
		panic("rotate error: parent cannot be nil")
	} else if parent == rbtree.root {
		return node
	}

	var grandparent *RbTreeNode[T] = parent.parent
	if grandparent == nil {
		// Assert grandparent node is not nil
		panic("rotate error: grandparent node should not be nil")
	}


	switch rotation {
	case LL:
		// Rotate and reassign parents
		grandparent.left = parent.right
		if parent.right != nil {
			parent.right.parent = grandparent
		}

		parent.right = grandparent
		parent.parent = grandparent.parent
		grandparent.parent = parent

		// Recolor
		grandparent.isBlack = false
		parent.isBlack = true

		return parent
	case LR:
		// Rotate and reassign parents
		parent.right = node.left
		if node.left != nil {
			node.left.parent = parent
		}

		node.left = parent
		parent.parent = node

		grandparent.left = node.right
		if node.right != nil {
			node.right.parent = grandparent
		}

		node.right = grandparent
		node.parent = grandparent.parent
		grandparent.parent = node

		// Recolor
		grandparent.isBlack = false
		node.isBlack = true

		return node
	case RR:
		// Rotate and reassign parents
		grandparent.right = parent.left
		if parent.left != nil {
			parent.left.parent = grandparent
		}

		parent.left = grandparent
		parent.parent = grandparent.parent
		grandparent.parent = parent

		// Recolor
		grandparent.isBlack = false
		parent.isBlack = true

		return parent
	case RL:
		// Rotate and reassign parents
		parent.left = node.right
		if node.right != nil {
			node.right.parent = parent
		}

		node.right = parent
		parent.parent = node

		grandparent.right = node.left
		if node.left != nil {
			node.left.parent = grandparent
		}

		node.left = grandparent
		node.parent = grandparent.parent
		grandparent.parent = node

		// Recolor
		grandparent.isBlack = false
		node.isBlack = true

		return node
	default:
		panic("Invalid rotation option provided")
	}
}

func (rbtree *RbTree[T]) recolor(curr *RbTreeNode[T]) {
	for curr != nil {
		if curr == rbtree.root {
			return
		}

		parent := curr.parent
		// TODO: remove assertion - Assert parent node is not nil
		if parent == nil {
			panic("Recolor error: parent node cannot be nil")
		}

		if parent == rbtree.root {
			// Parent node is tree root, no further recoloring needed
			return
		} else if curr.isBlack && parent.isBlack {
			// Current and parent nodes are black, no further recoloring needed
			return
		} else if curr.isBlack || parent.isBlack {
			// Current or parent nodes are black so current node should not be recolored, move to parent node
			curr = parent
			continue
		}

		grandparent := parent.parent
		// TODO: remove assertion - Assert grandparent node is not nil
		if grandparent == nil {
			panic("Recolor error: grandparent node cannot be nil")
		}

		var uncle *RbTreeNode[T]
		if grandparent.left == parent {
			uncle = grandparent.right
		} else if grandparent.right == parent {
			uncle = grandparent.left
		} else {
			//TODO: remove assertion - Assert uncle node reference is a valid child node
			panic("Recolor error: Invalid reference from grandparent node to child node")
		}

		greatgrandparent := grandparent.parent
		if greatgrandparent == nil && grandparent != rbtree.root {
			// TODO: remove assertion - Assert greatgrandparent must bil nil if grandparent is tree root
			panic("recolor error: greatgrandparent is nil while grandparent is not tree root")
		}

		if uncle != nil && !uncle.isBlack {
			// Case: uncle is red, recolor parent and uncle
			uncle.isBlack = true
			parent.isBlack = true
			// If grandparent is not root, recolor it red
			if greatgrandparent != nil {
				grandparent.isBlack = false
			} else {
				// TODO: Remove assertion - Assert grandparent node is root of tree
				if grandparent != rbtree.root {
					panic("recolor error: grandparent node is not the root of the tree")
				}
			}
			curr = parent
		} else {
			// Case: uncle is black or nil (empty black node), perform rotation
			// var temp *RbTreeNode[T]
			if grandparent.left == parent {
				// Case: Left leaning subtree, rotate right
				if parent.left == curr {
					// Perform LL rotation
					rbtree.rotate(curr, LL)
					// temp = parent
					curr = parent
				} else if parent.right == curr {
					// Perform LR rotation
					rbtree.rotate(curr, LR)
					// temp = curr
				} else {
					panic("Rotate error: Invalid left rotation scenario")
				}
			} else if grandparent.right == parent {
				// Case: Right leaning subtree, rotate left
				if parent.right == curr {
					// Perform RR rotation
					rbtree.rotate(curr, RR)
					// temp = parent
					curr = parent
				} else if parent.left == curr {
					// Perform RL rotation
					rbtree.rotate(curr, RL)
					// temp = curr
				} else {
					panic("Rotate error: Invalid right rotation scenario")
				}
			} else {
				panic("Rotate error: Invalid grandparent child node reference")
			}

			// Reassign greatgrandparent child reference to rotated sub tree or reassign tree root node
			if greatgrandparent == nil {
				// Case: Reassign tree root node
				// rbtree.root = temp
				rbtree.root = curr
				return
			} else if greatgrandparent.left == grandparent {
				// Case: Reassign left greatgrandparent child reference
				greatgrandparent.left = curr
				if curr != nil {
					curr.parent = greatgrandparent
				} else {
					// TODO: Remove assertion - Assert temp node is not nil
					panic("recolor error: rotated node should not be nil")
					// panic("recolor error: temp should not be nil")
				}
				// greatgrandparent.left = temp
				// if temp != nil {
				// 	temp.parent = greatgrandparent
				// } else {
				// 	// TODO: Remove assertion - Assert temp node is not nil
				// 	panic("recolor error: temp should not be nil")
				// }
			} else if greatgrandparent.right == grandparent {
				// Case: Reassign right greatgrandparent child reference
				greatgrandparent.right = curr
				if curr != nil {
					curr.parent = greatgrandparent
				} else {
					// TODO: Remove assertion - Assert temp node is not nil
					panic("recolor error: rotated node should not be nil")
					// panic("recolor error: temp should not be nil")
				}
				// greatgrandparent.right = temp
				// if temp != nil {
				// 	temp.parent = greatgrandparent
				// } else {
				// 	// TODO: Remove assertion - Assert temp node is not nil
				// 	panic("recolor error: temp should not be nil")
				// }
			} else {
				// TODO: remove assertion - Assert greatgrandparent child node references grandparent node
				panic("recolor error: Invalid reference from greatgrandparent node to grandparent node")
			}
			curr = greatgrandparent
		}
	}
}

func (rbtree *RbTree[T]) Remove(val T) {
	// TODO: Implement method
}

func (rbtree *RbTree[T]) IsValidRbTree() bool {
	if rbtree.capacity == 0 {
		// Determine if root node is nil if tree is empty
		return rbtree.root == nil
	} else if !rbtree.root.isBlack {
		// Root of tree is red violation
		return false
	}

	// Define local stack data object to store current node in DFS path and count of black nodes encountered
	type Level struct {
		node           *RbTreeNode[T]
		depth          int
		blackNodeCount int
	}

	// DFS for consecutive red node violations and black node path counts
	var counted uint = 0
	blackNodesInPath := 0
	maxTreeHeight := int(math.Log2(float64(rbtree.capacity))) + 1
	stack := make([]*Level, 1, maxTreeHeight)
	stack[0] = &Level{rbtree.root, 1, 1}
	for len(stack) > 0 {
		i := len(stack) - 1
		curr := stack[i].node
		depth := stack[i].depth
		count := stack[i].blackNodeCount
		counted++
		// Pop level from stack
		stack = stack[:i]

		// Assert node from current level is not nil as Levels containing nil nodes will never be pushed onto the stack
		if curr == nil {
			panic("Popped node from stack cannot be nil")
		}

		for _, child := range []*RbTreeNode[T]{curr.left, curr.right} {
			if child != nil {
				// Case: Child is not nil, check for consecutive red node violation or push next stack level
				if child.parent != curr {
					// Invalid reference from child to parent node
					return false
				} else if child.val == curr.val {
					// Two nodes exist with the same value
					return false
				} else if child.isBlack {
					// Case: current node is black, push next stack level and increment black node count
					stack = append(stack, &Level{child, depth + 1, count + 1})
				} else {
					// Case: current node is red
					if !curr.isBlack {
						// Two consecutive red nodes violation
						return false
					} else {
						stack = append(stack, &Level{child, depth + 1, count})
					}
				}
			} else {
				// Case: Child is nil, check terminal path black node count
				c := count + 1
				if blackNodesInPath == 0 {
					blackNodesInPath = c
				} else if blackNodesInPath != c {
					// Current terminal path does not contain the same number of black nodes as another terminal path violation
					return false
				}
			}
		}
	}

	if counted != rbtree.capacity {
		// Counted nodes in tree not equal to tree capacity
		return false
	}

	return true
}

func testInsert() {
	nums := []int{69, 1, 76, 3, 70, 72, 72, 4, 71, 71, 74, 73, 75}
	fmt.Println("Begin testInsert: ", nums)
	rbtree := CreateRbTree[int]()
	for _, num := range nums {
		fmt.Println("Inserting:", num)
		rbtree.Insert(num)
		rbtree.Print()
		if !rbtree.IsValidRbTree() {
			panic("Invalid red black tree structure")
		}
	}
	fmt.Println("End testInsert")
}

func main() {
	testInsert()
}
