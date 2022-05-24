package str

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, r := range t {
		m[r]++
	}

	for _, v := range s {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}

	return true
}
