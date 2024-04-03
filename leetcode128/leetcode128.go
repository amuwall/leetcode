package leetcode128

func longestConsecutive(nums []int) int {
	h := map[int]int{}
	for _, num := range nums {
		h[num] = 1
	}

	maxLen := 0
	for num, _ := range h {
		if h[num] == 0 {
			continue
		}

		i := num
		for {
			if _, ok := h[i+1]; !ok || h[i+1] == 0 {
				break
			}

			h[num] += h[i+1]
			h[i+1] = 0
			i++
		}

		if h[num] > maxLen {
			maxLen = h[num]
		}
	}

	return maxLen
}
