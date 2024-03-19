package leetcode392

func isSubsequence(s string, t string) bool {
	p := 0
	for i := 0; i < len(s); i++ {
		found := false
		for j := p; j < len(t); j++ {
			if s[i] == t[j] {
				found = true
				p = j + 1
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
