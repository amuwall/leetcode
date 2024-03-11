package leetcode2129

import "strings"

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	for i, word := range words {
		if len(word) < 3 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}

	return strings.Join(words, " ")
}
