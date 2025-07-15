package heap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	size   uint
	values []T
}

func CreateHeap[T constraints.Ordered](size uint) Heap[T] {
	return Heap[T]{values: make([]T, size)}
}

func Heapify[T constraints.Ordered](array []T) Heap[T] {
	n := len(array)
	heap := Heap[T]{size: uint(n), values: array}
	for i := len(array) - 1; i >= 0; i-- {
		heap.percolateUp(uint(i))
	}
	return heap
}

func (heap *Heap[T]) percolateUp(ind uint) {
	for ind > 0 {
		p_ind := ind / 2
		if heap.values[p_ind] > heap.values[ind] {
			tmp := heap.values[p_ind]
			heap.values[p_ind] = heap.values[ind]
			heap.values[ind] = tmp
			ind /= 2
		} else {
			return
		}
	}
}

func (heap *Heap[T]) percolateDown(ind uint) {
	for ind < heap.size {
		lc_ind, rc_ind := ind*2, ind*2+1
		var c_ind uint // Target child index
		if rc_ind < heap.size && heap.values[lc_ind] > heap.values[rc_ind] {
			c_ind = rc_ind
		} else {
			c_ind = lc_ind
		}

		if heap.values[c_ind] < heap.values[ind] {
			tmp := heap.values[c_ind]
			heap.values[c_ind] = heap.values[ind]
			heap.values[ind] = tmp
			ind *= 2
		} else {
			return
		}
	}
}

func (heap *Heap[T]) Insert(val T) bool {
	if heap.size > uint(len(heap.values)) {
		return false
	}
	heap.size++
	heap.values[heap.size-1] = val
	heap.percolateUp(heap.size - 1)
	return true
}

func (heap *Heap[T]) Extract() (val T, err error) {
	if heap.size < 1 {
		var missingVal T
		return missingVal, errors.New("Heap does not contain any values to be extracted")
	}
	ret := heap.values[0]
	heap.values[0] = heap.values[heap.size-1]
	heap.size--
	heap.percolateDown(0)
	return ret, nil
}

func (heap *Heap[T]) Clear() {
	heap.values = []T{}
	heap.size = 0
}
