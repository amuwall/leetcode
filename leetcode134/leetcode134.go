package leetcode134

func canCompleteCircuit(gas []int, cost []int) int {
	totalSum := 0
	currentSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		v := gas[i] - cost[i]
		totalSum += v
		currentSum += v
		if currentSum < 0 {
			currentSum = 0
			start = i + 1
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
