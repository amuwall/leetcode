package leetcode2024

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxConsecutiveChar(answerKey string, k int, ch byte) int {
	head := 0
	sumOfChange := 0

	maxLength := 0

	for tail := range answerKey {
		if answerKey[tail] != ch {
			sumOfChange++
		}
		for sumOfChange > k {
			if answerKey[head] != ch {
				sumOfChange--
			}

			head++
		}

		maxLength = max(maxLength, tail-head+1)
	}

	return maxLength
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(
		maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'),
	)
}
