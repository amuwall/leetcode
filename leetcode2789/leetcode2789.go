package leetcode2789

func maxArrayValue(nums []int) int64 {
	result := int64(nums[len(nums)-1])

	for i := len(nums) - 2; i >= 0; i-- {
		if result >= int64(nums[i]) {
			result += int64(nums[i])
		} else {
			result = int64(nums[i])
		}
	}

	return result
}
