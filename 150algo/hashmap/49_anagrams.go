package hashmap

import "sort"

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)

	for _, s := range strs {
		runeSlice := []rune(s)
		sort.Slice(runeSlice, func(i int, j int) bool {
			return runeSlice[i] < runeSlice[j]
		})

		m[string(runeSlice)] = append(m[string(runeSlice)], s)
	}

	result := make([][]string, len(m))
	i := 0
	for _, v := range m {
		result[i] = v
		i++
	}
	return result
}
