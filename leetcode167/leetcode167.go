package leetcode167

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
