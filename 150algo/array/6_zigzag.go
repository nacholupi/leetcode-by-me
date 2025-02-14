package array

import "strings"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	var result strings.Builder

	cycleLen := (numRows - 1) * 2

	for i := 0; i < numRows; i++ {
		pos := i
		middleRow := i != 0 && i != numRows-1
		for pos < len(s) {
			result.WriteByte(s[pos])
			nextPos := pos + cycleLen - 2*i
			if middleRow && nextPos < len(s) {
				result.WriteByte(s[nextPos])
			}
			pos += cycleLen
		}
	}
	return result.String()
}
