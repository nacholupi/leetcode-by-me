package array

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			} else if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}
