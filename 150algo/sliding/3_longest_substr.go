package sliding

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 1
	dis := 1

	for right < len(s) {
		idx := strings.Index(s[left:right], string(s[right]))
		if idx != -1 {
			left += idx + 1
		}

		right++

		if dis < right-left {
			dis = right - left
		}
	}

	return dis
}
