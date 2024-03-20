package leetcode151

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
