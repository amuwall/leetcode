package leetcode2529

func maximumCount(nums []int) int {
	return max(lowerBound(nums, 0), len(nums)-lowerBound(nums, 1))
}

func lowerBound(nums []int, v int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) >> 1
		if nums[m] >= v {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
