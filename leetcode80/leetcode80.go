package leetcode80

func removeDuplicates(nums []int) int {
	if len(nums) == 1 || len(nums) == 2 {
		return len(nums)
	}

	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[j-1] && nums[i] == nums[j-2] {
			continue
		}

		nums[j] = nums[i]
		j++
	}

	return j
}
