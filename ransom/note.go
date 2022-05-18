package ransom

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) <= 0 || len(magazine) <= 0 ||
		len(magazine) < len(ransomNote) {
		return false
	}

	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}

	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
