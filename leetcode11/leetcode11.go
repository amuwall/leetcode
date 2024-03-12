package leetcode11

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	result := 0
	for {
		if i >= j {
			break
		}

		result = max(result, min(height[i], height[j])*(j-i))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
