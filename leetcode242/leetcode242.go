package leetcode242

func isAnagram(s string, t string) bool {
	chars := map[int32]int{}
	for _, char := range s {
		if _, ok := chars[char]; ok {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}

	for _, char := range t {
		if _, ok := chars[char]; ok {
			chars[char]--
			if chars[char] == 0 {
				delete(chars, char)
			}
		} else {
			return false
		}
	}

	return len(chars) == 0
}
