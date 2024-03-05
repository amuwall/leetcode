package leetcode58

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result == 0 {
				continue
			} else {
				break
			}
		}

		result++
	}

	return result
}
