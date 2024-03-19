package leetcode45

func jump(nums []int) int {
	steps := 0
	maxPosition := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
