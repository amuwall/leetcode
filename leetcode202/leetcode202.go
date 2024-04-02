package leetcode202

func isHappy(n int) bool {
	nums := map[int]bool{}

	for {
		if _, ok := nums[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}

		nums[n] = true

		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		n = sum
	}
}
