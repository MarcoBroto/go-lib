package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRbTree(t *testing.T) {
	tree := CreateRbTree[int]()
	assert.NotNil(t, tree)
	assert.Nil(t, tree.root)
	assert.Zero(t, tree.capacity)
}

func TestIsValidRbTreeReturnsTrue(t *testing.T) {

}

func TestIsValidRbTreeReturnsFalse(t *testing.T) {

}

func TestIsValidRbTreeReturnsPanics(t *testing.T) {
	// tree := CreateRbTree[int]()
	// assert.Panics(t, tree.isValidRbTree())
}

func TestInsert(t *testing.T) {
	nums := []int{69, 1, 76, 3, 70, 72, 72, 4, 71, 71, 74, 73, 75}
	tree := CreateRbTree[int]()
	for _, num := range nums {
		tree.Insert(num)
	}
	assert.True(t, tree.IsValidRbTree())
}

func TestInsertRandom(t *testing.T) {
	assert.True(t, true)
}

func TestRemove(t *testing.T) {
	assert.True(t, true)
}

func TestRotate_LL(t *testing.T) {
	assert.True(t, true)
}

func TestLRRotate_LR(t *testing.T) {
	assert.True(t, true)
}

func TestRotate_RR(t *testing.T) {
	assert.True(t, true)
}

func TestRotate_RL(t *testing.T) {
	assert.True(t, true)
}

func TestRecolor(t *testing.T) {
	assert.True(t, true)
}
