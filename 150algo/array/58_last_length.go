package array

func lengthOfLastWord(s string) int {
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			j++
		}
		if s[i] == ' ' && j != 0 {
			break
		}
	}
	return j
}
