package leetcode14

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	s := strs[0]
	prefix := s
	for i := len(s); i >= 0; i-- {
		prefix = s[:i]
		allHas := true
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], prefix) {
				allHas = false
				break
			}
		}
		if allHas {
			break
		}
	}

	return prefix
}
