package intervals

import "sort"

func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	if len(intervals) == 0 {
		return result
	}
	ci := []int{intervals[0][0], intervals[0][1]}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ci[1] {
			if intervals[i][1] < ci[1] {
				continue
			} else {
				ci[1] = intervals[i][1]
			}
		} else {
			result = append(result, ci)
			ci = []int{intervals[i][0], intervals[i][1]}
		}
	}
	return append(result, ci)
}
