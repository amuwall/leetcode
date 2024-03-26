package leetcode135

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	result := max(1, left[len(left)-1])
	preRight := 1
	for i := len(ratings) - 2; i >= 0; i-- {
		right := 1
		if ratings[i] > ratings[i+1] {
			right = preRight + 1
		}

		result += max(right, left[i])

		preRight = right
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
