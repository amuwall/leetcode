package leetcode68

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	line := ""
	for _, word := range words {
		if len(line) == 0 {
			line = word
			continue
		}
		if len(line)+len(word)+1 > maxWidth {
			lines = append(lines, line)
			line = word
		} else {
			line += " " + word
		}
	}
	lines = append(lines, line)

	result := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		lineWords := strings.Split(lines[i], " ")
		if len(lineWords) == 1 || i == len(lines)-1 {
			result[i] = lines[i] + strings.Repeat(" ", maxWidth-len(lines[i]))
			continue
		}

		numSpaces := maxWidth - len(lines[i]) + len(lineWords) - 1
		avgSpaces := numSpaces / (len(lineWords) - 1)
		extraSpaces := numSpaces % (len(lineWords) - 1)

		result[i] = lineWords[0]
		for j := 1; j < len(lineWords); j++ {
			result[i] = result[i] + strings.Repeat(" ", avgSpaces)
			if extraSpaces > 0 {
				result[i] += " "
				extraSpaces--
			}
			result[i] += lineWords[j]
		}
	}

	return result
}
