package main

func partitionString(s string) int {
	i := 0
	substring := ""
	for _, runeInput := range s {

		for _, runeSubstr := range substring {
			if runeSubstr == runeInput {
				substring = ""
				break
			}
		}

		if substring == "" {
			i++
		}
		substring += string(runeInput)
	}
	return i
}
