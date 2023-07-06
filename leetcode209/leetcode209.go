package leetcode209

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	head := 0
	tail := 0

	minLength := 0

	for _, num := range nums {
		sum = sum + num
		tail += 1

		for {
			if sum >= target {
				if tail-head <= minLength || minLength == 0 {
					minLength = tail - head
				}

				sum = sum - nums[head]
				head += 1
			} else {
				break
			}
		}
	}

	return minLength
}
