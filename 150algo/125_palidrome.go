package twopointers

import (
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		a := unicode.ToLower(rune(s[i]))
		b := unicode.ToLower(rune(s[j]))
		if !(unicode.IsLetter(a) || unicode.IsDigit(a)) {
			i++
			continue
		}
		if !(unicode.IsLetter(b) || unicode.IsDigit(b)) {
			j--
			continue
		}

		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}
