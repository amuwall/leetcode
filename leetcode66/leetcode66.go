package leetcode66

func plusOne(digits []int) []int {
	plus := false
	result := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		result[i] = digits[i]
		if i == len(digits)-1 {
			result[i] += 1
		}
		if plus {
			result[i] += 1
			plus = false
		}
		if result[i] >= 10 {
			result[i] = result[i] % 10
			plus = true
		}
	}

	if plus {
		result = append([]int{1}, result...)
	}

	return result
}
