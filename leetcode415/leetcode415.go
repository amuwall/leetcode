package leetcode415

func addStrings(num1 string, num2 string) string {
	result := ""

	i := len(num1) - 1
	j := len(num2) - 1
	plus := false
	for {
		if i < 0 && j < 0 {
			break
		}

		var val uint8
		if i >= 0 {
			val += num1[i] - '0'
		}
		if j >= 0 {
			val += num2[j] - '0'
		}
		if plus {
			val += 1
			plus = false
		}
		if val >= 10 {
			plus = true
			val -= 10
		}

		result = string(val+'0') + result

		i--
		j--
	}

	if plus {
		result = "1" + result
	}

	return result
}
