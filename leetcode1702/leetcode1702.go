package leetcode1702

func maximumBinaryString(binary string) string {
	n := len(binary)
	result := []rune(binary)

	j := 0
	for i := 0; i < n; i++ {
		if result[i] == '0' {
			for j <= i || (j < n && result[j] == '1') {
				j++
			}

			if j < n {
				result[i] = '1'
				result[j] = '1'
				result[i+1] = '0'
			}
		}
	}

	return string(result)
}
