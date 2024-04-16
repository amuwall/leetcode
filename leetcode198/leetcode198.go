package leetcode198

func rob(nums []int) int {
	sum := make([][2]int, len(nums))

	sum[0][0] = 0
	sum[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i][0] = max(sum[i-1][0], sum[i-1][1])
		sum[i][1] = sum[i-1][0] + nums[i]
	}

	return max(sum[len(nums)-1][0], sum[len(nums)-1][1])
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
