package bit

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var res string
	carry := 0
	lenA := len(a)
	lenB := len(b)
	maxLen := lenA
	if lenA < lenB {
		maxLen = lenB
	}

	for i := 0; i < maxLen; i++ {
		aDigit := 0
		bDigit := 0
		if i < len(a) && a[len(a)-1-i] == '1' {
			aDigit = 1
		}
		if i < len(b) && b[len(b)-1-i] == '1' {
			bDigit = 1
		}

		add := carry + aDigit + bDigit
		mod := add % 2
		carry = int(add / 2)
		res = fmt.Sprintf("%d%s", mod, res)
	}

	if carry == 1 {
		res = fmt.Sprintf("1%s", res)
	}

	return res
}
