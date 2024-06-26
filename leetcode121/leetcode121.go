package leetcode121

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for i := 1; i <= len(prices)-1; i++ {
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return profit
}
