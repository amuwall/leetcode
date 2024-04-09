package leetcode35

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
