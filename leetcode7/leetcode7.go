package leetcode7

import "math"

func reverse(x int) int {
	result := 0
	for {
		if x == 0 {
			break
		}

		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			result = 0
			break
		}

		num := x % 10
		x = x / 10
		result = result*10 + num
	}

	return result
}
