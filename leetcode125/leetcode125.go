package leetcode125

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			break
		}

		if !isalnum(s[i]) {
			i++
			continue
		}

		if !isalnum(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
