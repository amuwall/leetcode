package leetcode1793

func maximumScore(nums []int, k int) int {
	minNums := make([]int, len(nums))
	minNums[k] = nums[k]
	for i := k - 1; i >= 0; i-- {
		minNums[i] = min(nums[i], minNums[i+1])
	}
	for j := k + 1; j < len(nums); j++ {
		minNums[j] = min(nums[j], minNums[j-1])
	}

	result := nums[k]
	i := k
	j := k
	for i >= 0 && j < len(nums) {
		for i > 0 && minNums[i-1] == minNums[i] {
			i--
		}
		for j < len(nums)-1 && minNums[j+1] == minNums[j] {
			j++
		}

		result = max(result, min(minNums[i], minNums[j])*(j-i+1))
		if i > 0 && j < len(nums)-1 {
			if minNums[i-1] < minNums[j+1] {
				j++
			} else {
				i--
			}
		} else if i > 0 {
			i--
		} else {
			j++
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
