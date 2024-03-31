package leetcode12

var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romanInt = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func intToRoman(num int) string {
	result := ""
	for num != 0 {
		for i := len(romanInt) - 1; i >= 0; i-- {
			if num >= romanInt[i] {
				result += romanMap[romanInt[i]]
				num -= romanInt[i]
				break
			}
		}
	}

	return result
}
