package leetcode2908

func minimumSum(nums []int) int {
	leftMinNum := make([]int, len(nums))
	rightMinNum := make([]int, len(nums))

	leftMinNum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftMinNum[i] = min(leftMinNum[i-1], nums[i])
	}
	rightMinNum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightMinNum[i] = min(rightMinNum[i+1], nums[i])
	}

	result := -1
	for i := 1; i < len(nums)-1; i++ {
		if leftMinNum[i-1] >= nums[i] || rightMinNum[i+1] >= nums[i] {
			continue
		}

		if result == -1 {
			result = leftMinNum[i-1] + nums[i] + rightMinNum[i+1]
		} else {
			result = min(result, leftMinNum[i-1]+nums[i]+rightMinNum[i+1])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
