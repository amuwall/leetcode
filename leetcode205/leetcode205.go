package leetcode205

func isIsomorphic(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := s2t[s[i]]; ok {
			if s2t[s[i]] != t[i] {
				return false
			}
		} else {
			if _, ok := t2s[t[i]]; ok {
				return false
			}

			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		}
	}

	return true
}
