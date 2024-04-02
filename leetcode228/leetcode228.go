package leetcode228

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	ranges := []string{}
	if len(nums) == 0 {
		return ranges
	}

	start := 0
	end := 0
	for end = 1; end < len(nums); end++ {
		if nums[end]-nums[end-1] == 1 {
			continue
		}

		if start+1 == end {
			ranges = append(ranges, fmt.Sprintf("%d", nums[start]))
		} else {
			ranges = append(ranges, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
		}

		start = end
	}

	if start+1 == end {
		ranges = append(ranges, fmt.Sprintf("%d", nums[start]))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
	}

	return ranges
}
