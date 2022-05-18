package firstunique

func firstUniqChar(s string) int {

	if len(s) == 0 {
		return -1
	}

	if len(s) == 1 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		found := false
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}

			if s[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}

	return -1
}
