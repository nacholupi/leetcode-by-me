package hashmap

func isIsomorphic(s string, t string) bool {
	stMap := make(map[byte]byte)
	tsMap := make(map[byte]byte)

	if s == "" && t == "" {
		return true
	}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if tString, ok := stMap[s[i]]; ok {
			if t[i] != tString {
				return false
			}
		} else {
			if _, ok := tsMap[t[i]]; ok {
				return false
			}
			stMap[s[i]] = t[i]
			tsMap[t[i]] = s[i]
		}
	}
	return true
}
