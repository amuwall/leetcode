package leetcode2391

import "strings"

func garbageCollection(garbage []string, travel []int) int {
	total := 0

	for _, i := range travel {
		total += i * 3
	}

	for _, s := range garbage {
		total += len(s)
	}

	for i := len(garbage) - 1; i > 0; i-- {
		if !strings.Contains(garbage[i], "M") {
			total -= travel[i-1]
		} else {
			break
		}
	}

	for i := len(garbage) - 1; i > 0; i-- {
		if !strings.Contains(garbage[i], "P") {
			total -= travel[i-1]
		} else {
			break
		}
	}

	for i := len(garbage) - 1; i > 0; i-- {
		if !strings.Contains(garbage[i], "G") {
			total -= travel[i-1]
		} else {
			break
		}
	}

	return total
}
