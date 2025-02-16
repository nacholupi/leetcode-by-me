package sliding

func minWindow(s string, t string) string {

	result := ""

	freq := make(map[rune]int)
	for _, c := range t {
		freq[c]++
	}

	l, r, filled := 0, 0, 0

	for r < len(s) {
		cr := rune(s[r])
		cfv, exists := freq[cr]
		if !exists {
			r++
			continue
		}

		if cfv == 1 {
			filled++
		}
		freq[cr]--

		if filled != len(freq) {
			r++
			continue
		}

		r++
		if len(result) == 0 || len(result) > r-l {
			result = s[l:r]
		}

		for l < r {
			cl := rune(s[l])
			cfv, exists := freq[cl]
			if exists {
				freq[cl]++
				if cfv == 0 {
					filled--
				}
				if filled != len(freq) {
					l++
					break
				}
			}
			l++
			if len(result) > r-l {
				result = s[l:r]
			}
		}
	}

	return result
}
