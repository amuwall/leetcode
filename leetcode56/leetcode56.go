package leetcode56

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{
		intervals[0],
	}
	k := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merged[k][1] {
			if intervals[i][1] > merged[k][1] {
				merged[k][1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
			k++
		}
	}

	return merged
}
