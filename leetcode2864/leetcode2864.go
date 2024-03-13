package leetcode2864

func maximumOddBinaryNumber(s string) string {
	numsOfZero := 0
	numsOfOne := 0

	for _, ch := range s {
		if ch == '0' {
			numsOfZero++
		} else {
			numsOfOne++
		}
	}

	result := ""
	for i := 0; i < numsOfOne-1; i++ {
		result += "1"
	}
	for i := 0; i < numsOfZero; i++ {
		result += "0"
	}
	if numsOfOne > 0 {
		result += "1"
	}

	return result
}
