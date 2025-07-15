package forest

func Find(forest []int, t int) int {
	if forest[t] < 0 {
		return t
	} else {
		forest[t] = Find(forest, forest[t])
		return forest[t]
	}
}

func Union(forest []int, u int, v int) {
	u, v = Find(forest, u), Find(forest, v)

	if forest[u] <= forest[v] {
		forest[v] = u
		forest[u]--
	} else {
		forest[u] = v
		forest[v]--
	}
}

func CreateForest(size uint) []int {
	forest := make([]int, size)
	for ind := range forest {
		forest[ind] = -1
	}
	return forest
}
