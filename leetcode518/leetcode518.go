package leetcode518

func change(amount int, coins []int) int {
	counts := make([]int, amount+1)
	counts[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if counts[i-coin] != 0 {
				counts[i] = counts[i] + counts[i-coin]
			}
		}
	}

	return counts[amount]
}
