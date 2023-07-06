package leetcode1480

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			result[i] = num
		} else {
			result[i] = result[i-1] + num
		}
	}

	return result
}
