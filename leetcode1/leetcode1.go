package leetcode1

func twoSum(nums []int, target int) []int {
	existNums := map[int]int{}

	for i, num := range nums {
		needNum := target - num

		j, ok := existNums[needNum]
		if ok {
			return []int{j, i}
		}

		existNums[num] = i
	}

	return []int{}
}
