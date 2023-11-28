package leetcode9

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	half := 0
	for x > half {
		half = half*10 + x%10
		x = x / 10
	}

	return half == x || x == half/10
}
