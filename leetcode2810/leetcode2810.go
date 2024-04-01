package leetcode2810

import (
	"slices"
)

func finalString(s string) string {
	chars := []uint8{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			slices.Reverse(chars)
		} else {
			chars = append(chars, s[i])
		}
	}

	result := ""
	for _, char := range chars {
		result += string(char)
	}

	return result
}
