package leetcode139

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}

			for _, word := range wordDict {
				if s[j:i] == word {
					dp[i] = true
					break
				}
			}

			if dp[i] {
				break
			}
		}
	}

	return dp[len(s)]
}
