package leetcode3

func lengthOfLongestSubstring(s string) int {
	existCharsIndex := map[uint8]int{}

	maxLength := 0
	leftIndex := 0
	rightIndex := 0

	for rightIndex = 0; rightIndex < len(s); rightIndex++ {
		existCharIndex, ok := existCharsIndex[s[rightIndex]]
		if ok && leftIndex <= existCharIndex {
			leftIndex = existCharIndex + 1
		} else {
			existCharsIndex[s[rightIndex]] = rightIndex
		}

		if rightIndex-leftIndex+1 > maxLength {
			maxLength = rightIndex - leftIndex + 1
		}

		existCharsIndex[s[rightIndex]] = rightIndex
	}

	return maxLength
}
