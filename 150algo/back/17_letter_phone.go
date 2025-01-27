package back

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	letters := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var result []string

	for _, c := range letters[digits[0]] {
		result = append(result, string(c))
	}

	for i := 1; i < len(digits); i++ {
		d := digits[i]
		var dResult []string
		for _, c := range letters[d] {
			for _, re := range result {
				dResult = append(dResult, re+string(c))
			}
		}
		result = dResult
	}
	return result
}
