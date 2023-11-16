package leetcode1535

func getWinner(arr []int, k int) int {
	winner := arr[0]
	winCount := 0

	for _, num := range arr[1:] {
		if winner > num {
			winCount++
		} else {
			winner = num
			winCount = 1
		}

		if winCount == k {
			break
		}
	}

	return winner
}
