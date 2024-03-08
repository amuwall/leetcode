package leetcode2834

func minimumPossibleSum(n int, target int) int {
	const mod = 1000000007
	m := target / 2
	if n <= m {
		return (n * (n + 1) / 2) % mod
	}
	return ((m * (m + 1) / 2) + ((target + target + (n - m) - 1) * (n - m) / 2)) % mod
}
