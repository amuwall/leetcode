package leetcode274

import "slices"

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(a, b int) int { return b - a })

	result := 0
	for _, citation := range citations {
		if citation <= result {
			break
		}
		result++
	}

	return result
}
