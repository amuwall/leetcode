package leetcode42

func trap(height []int) int {
	leftMaxHeight := make([]int, len(height))
	rightMaxHeight := make([]int, len(height))

	leftMaxHeight[0] = 0
	for i := 1; i < len(height); i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i-1])
	}
	rightMaxHeight[len(height)-1] = 0
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i+1])
	}

	result := 0
	for i := 0; i < len(height); i++ {
		h := min(leftMaxHeight[i], rightMaxHeight[i])
		if h < height[i] {
			continue
		}

		result += h - height[i]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
