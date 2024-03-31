package leetcode331

import "strings"

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")

	slot := 1
	for _, node := range nodes {
		if slot == 0 {
			return false
		}

		if node == "#" {
			slot--
		} else {
			slot++
		}
	}

	return slot == 0
}
