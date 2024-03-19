package leetcode6

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	rowIndex := 0
	down := true
	for i := 0; i < len(s); i++ {
		rows[rowIndex] = rows[rowIndex] + s[i:i+1]

		if down {
			rowIndex++
			if rowIndex == numRows {
				down = false
				rowIndex = numRows - 2
			}
		} else {
			rowIndex--
			if rowIndex < 0 {
				down = true
				rowIndex = 1
			}
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
