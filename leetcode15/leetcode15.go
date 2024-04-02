package leetcode15

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := len(nums) - 1
		target := -nums[first]
		for second := first + 1; second < len(nums)-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return result
}
