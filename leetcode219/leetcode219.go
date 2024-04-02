package leetcode219

func containsNearbyDuplicate(nums []int, k int) bool {
	h := map[int][]int{}
	for i, num := range nums {
		h[num] = append(h[num], i)
	}

	for _, index := range h {
		if len(index) > 1 {
			for i := 1; i < len(index); i++ {
				if index[i]-index[i-1] <= k {
					return true
				}
			}
		}
	}

	return false
}
