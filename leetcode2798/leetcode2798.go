package leetcode2798

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	sum := 0
	for _, hour := range hours {
		if hour >= target {
			sum++
		}
	}

	return sum
}
