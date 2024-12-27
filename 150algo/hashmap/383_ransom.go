package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	magMap := map[rune]int{}
	for _, char := range magazine {
		magMap[char]++
	}

	for _, ransomChar := range ransomNote {
		if charCounter, ok := magMap[ransomChar]; ok && charCounter != 0 {
			magMap[ransomChar]--
			continue
		}
		return false
	}
	return true
}
