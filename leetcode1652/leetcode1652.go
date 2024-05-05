package leetcode1652

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	if k == 0 {
		return result
	}

	for i := 0; i < len(code); i++ {
		if k > 0 {
			count := 0
			j := (i + 1) % len(code)
			for count < k {
				result[i] += code[j]
				count++
				j = (j + 1) % len(code)
			}
		} else {
			count := 0
			j := (i - 1 + len(code)) % len(code)
			for count > k {
				result[i] += code[j]
				count--
				j = (j - 1 + len(code)) % len(code)
			}
		}
	}

	return result
}
