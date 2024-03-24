package leetcode322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sum := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		sum[i] = -1
		for _, coin := range coins {
			if i < coin {
				continue
			}
			if i == coin {
				sum[i] = 1
			}

			if sum[i-coin] == -1 {
				continue
			}
			if sum[i] == -1 {
				sum[i] = sum[i-coin] + 1
				continue
			}
			sum[i] = min(sum[i], sum[i-coin]+1)
		}
	}

	return sum[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
