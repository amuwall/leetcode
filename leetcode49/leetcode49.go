package leetcode49

func groupAnagrams(strs []string) [][]string {
	charsStatics := make([]map[byte]int, len(strs))
	groups := map[int][]string{}
	for i, str := range strs {
		chars := map[byte]int{}
		for j := 0; j < len(str); j++ {
			chars[str[j]]++
		}

		charsStatics[i] = chars
		k := i

		for j := 0; j < i; j++ {
			if equalMap(charsStatics[j], charsStatics[i]) {
				k = j
				break
			}
		}

		groups[k] = append(groups[k], str)
	}

	results := [][]string{}
	for _, group := range groups {
		results = append(results, group)
	}

	return results
}

func equalMap(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
