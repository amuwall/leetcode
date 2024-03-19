package leetcode55

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	step := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= step {
			step = 1
		} else {
			step++
		}
	}

	return step == 1
}
