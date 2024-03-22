package leetcode383

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	magazineChars := make(map[rune]int)
	for _, ch := range magazine {
		if _, ok := magazineChars[ch]; !ok {
			magazineChars[ch] = 1
		} else {
			magazineChars[ch]++
		}
	}

	for _, ch := range ransomNote {
		if _, ok := magazineChars[ch]; !ok {
			return false
		}

		magazineChars[ch]--
		if magazineChars[ch] < 0 {
			return false
		}
	}

	return true
}
