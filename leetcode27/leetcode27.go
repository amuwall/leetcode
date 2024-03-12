package leetcode27

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for {
		if i > j {
			break
		}

		if nums[i] == val {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp

			j--
		} else {
			i++
		}
	}

	return i
}
