package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walk(t, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		} else if ok1 != ok2 || v1 != v2 { // alt: !ok1 || !ok2 || v1 != v2
			return false
		}
	}
}

func TestEqualBinaryTree() { // main
	fmt.Println(Same(tree.New(1), tree.New(1)),
		Same(tree.New(1), tree.New(2)))
}
