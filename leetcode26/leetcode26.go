package leetcode26

func removeDuplicates(nums []int) int {
	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		nums[j] = nums[i]
		j++
	}

	return j
}
