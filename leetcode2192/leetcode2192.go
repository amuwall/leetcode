package leetcode2192

import "slices"

func getAncestors(n int, edges [][]int) [][]int {
	ancestors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		ancestors[i] = map[int]bool{}
	}

	e := make([][]int, n)
	indeg := make([]int, n)
	for _, edge := range edges {
		e[edge[0]] = append(e[edge[0]], edge[1])
		indeg[edge[1]]++
	}

	var q []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range e[u] {
			ancestors[v][u] = true
			for k := range ancestors[u] {
				ancestors[v][k] = true
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = []int{}
		for k := range ancestors[i] {
			result[i] = append(result[i], k)
		}
		slices.Sort(result[i])
	}

	return result
}
