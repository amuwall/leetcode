package leetcode1491

func average(salary []int) float64 {
	sum := 0
	minSalary := salary[0]
	maxSalary := salary[0]

	for _, v := range salary {
		sum += v
		if v < minSalary {
			minSalary = v
		}
		if v > maxSalary {
			maxSalary = v
		}
	}

	sum -= maxSalary
	sum -= minSalary

	return float64(sum) / float64(len(salary)-2)
}
