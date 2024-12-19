package array

func longestCommonPrefix(strs []string) string {

	result := ""
	shouldExit := false
	for charIdx := range strs[0] {
		for strIdx := 1; strIdx < len(strs); strIdx++ {
			if charIdx >= len(strs[strIdx]) || strs[0][charIdx] != strs[strIdx][charIdx] {
				shouldExit = true
				break
			}
		}
		if shouldExit {
			break
		}
		result += string(strs[0][charIdx])
	}

	return result
}
