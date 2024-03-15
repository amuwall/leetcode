package leetcode70

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	count := make([]int, n)
	count[0] = 1
	count[1] = 2
	for i := 2; i < n; i++ {
		count[i] += count[i-1] + count[i-2]
	}

	return count[n-1]
}
