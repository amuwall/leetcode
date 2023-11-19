package leetcode1743

func restoreArray(adjacentPairs [][]int) []int {
	pairs := map[int][]int{}
	for _, adjacentPair := range adjacentPairs {
		x := adjacentPair[0]
		y := adjacentPair[1]
		if _, ok := pairs[x]; ok {
			pairs[x] = append(pairs[x], y)
		} else {
			pairs[x] = []int{y}
		}
		if _, ok := pairs[y]; ok {
			pairs[y] = append(pairs[y], x)
		} else {
			pairs[y] = []int{x}
		}
	}

	var start int
	for num, _ := range pairs {
		if len(pairs[num]) == 1 {
			start = num
			break
		}
	}

	result := []int{start}
	var left int
	current := start
	for i := 0; i < len(adjacentPairs); i++ {
		pair := pairs[current]

		var nextValue int
		if pair[0] != left {
			nextValue = pair[0]
		} else {
			nextValue = pair[1]
		}

		result = append(result, nextValue)
		left = current
		current = nextValue
	}

	return result
}
