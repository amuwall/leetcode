package leetcode290

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	ch2word := map[rune]string{}
	word2ch := map[string]rune{}
	for i, ch := range pattern {
		if _, ok := ch2word[ch]; ok {
			if ch2word[ch] != words[i] {
				return false
			}
		} else {
			if _, ok := word2ch[words[i]]; ok {
				return false
			}
		}

		ch2word[ch] = words[i]
		word2ch[words[i]] = ch
	}

	return true
}
