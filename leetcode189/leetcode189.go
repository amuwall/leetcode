package leetcode189

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, num := range nums {
		newNums[(i+k)%len(nums)] = num
	}

	copy(nums, newNums)
}
