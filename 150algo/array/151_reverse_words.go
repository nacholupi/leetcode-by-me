package array

import (
	"strings"
)

func reverseWords(s string) string {
	ssP := strings.Split(s, " ")

	ss := []string{}
	for i := len(ssP) - 1; i >= 0; i-- {
		if len(ssP[i]) > 0 {
			ss = append(ss, ssP[i])
		}
	}

	return strings.Join(ss, " ")
}
