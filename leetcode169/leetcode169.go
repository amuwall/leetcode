package leetcode169

func majorityElement(nums []int) int {
	key := nums[0]
	count := 0

	for _, num := range nums {
		if num == key {
			count++
		} else {
			count--
		}

		if count == 0 {
			key = num
			count = 1
		}
	}

	return key
}
