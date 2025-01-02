package hashmap

import "strings"

func wordPattern(pattern string, s string) bool {

	patternMap := make(map[byte]string)
	stringMap := make(map[string]byte)

	if pattern == "" && s == "" {
		return true
	}

	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		lp := pattern[i]
		if mappedString, ok := patternMap[lp]; ok {
			if ss[i] != mappedString {
				return false
			}
		} else {
			if _, ok := stringMap[ss[i]]; ok {
				return false
			}
			patternMap[lp] = ss[i]
			stringMap[ss[i]] = lp
		}
	}
	return true
}
