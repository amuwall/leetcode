package leetcode238

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	result := make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = right[i]
		}
		if i == len(nums)-1 {
			result[i] = left[i]
		}
		result[i] = left[i] * right[i]
	}

	return result
}
