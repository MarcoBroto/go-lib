package forest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForestUnion(t *testing.T) {
	forest := []int{-1, -1, -1, -1, -1, -1}
	Union(forest, 3, 4)
	assert.Equal(t, forest, []int{-1, -1, -1, -2, 3, -1})
}

func TestForestFind(t *testing.T) {
	forest := []int{1, 2, 5, -1, -1, -4}
	assert.Equal(t, Find(forest, 3), 3)
	assert.Equal(t, Find(forest, 1), 5)
	assert.Equal(t, forest[5], -4)
	assert.Equal(t, Find(forest, 2), 5)
	assert.Equal(t, Find(forest, 5), 5)
	assert.Equal(t, forest, []int{1, 5, 5, -1, -1, -4})
	assert.Equal(t, Find(forest, 0), 5)
	assert.Equal(t, forest, []int{5, 5, 5, -1, -1, -4})
}

func TestCreateForest(t *testing.T) {
	forest := CreateForest(6)
	assert.Equal(t, forest, []int{-1, -1, -1, -1, -1, -1})
	forest = CreateForest(0)
	assert.Equal(t, forest, []int{})
	assert.NotEqual(t, forest, []int{-1})
}
