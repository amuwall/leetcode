package leetcode1759

func countHomogenous(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 1
	subStringLen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			subStringLen = (subStringLen + 1) % 1000000007
		} else {
			subStringLen = 1
		}
		count = (count + subStringLen) % 1000000007
	}

	return count
}
