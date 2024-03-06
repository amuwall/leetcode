package leetcode2917

func findKOr(nums []int, k int) int {
	result := 0

	for i := 0; i < 31; i++ {
		count := 0
		for _, num := range nums {
			count += (num >> i) & 1
		}
		if count >= k {
			result |= 1 << i
		} else {
			result |= 0
		}
	}

	return result
}
