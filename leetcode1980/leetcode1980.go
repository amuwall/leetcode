package leetcode1980

func findDifferentBinaryString(nums []string) string {
	s := ""
	for i, num := range nums {
		if num[i] == '0' {
			s += "1"
		} else {
			s += "0"
		}
	}

	return s
}
