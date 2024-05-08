package leetcode2079

func wateringPlants(plants []int, capacity int) int {
	steps := 0

	currentWater := capacity
	for i, needWater := range plants {
		if currentWater < needWater {
			currentWater = capacity
			steps += i + i
		}

		steps++
		currentWater -= needWater
	}

	return steps
}
