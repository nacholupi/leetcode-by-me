package array

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, v := range citations {
		if i >= v {
			return i
		}
	}
	return len(citations)
}
