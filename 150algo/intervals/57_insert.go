package intervals

import "math"

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	newLo := newInterval[0]
	newHi := newInterval[1]

	if intervals[len(intervals)-1][1] < newLo {
		return append(intervals, newInterval)
	}

	result := make([][]int, 0)
	i := 0

	for i < len(intervals) {
		intLo := intervals[i][0]
		intHi := intervals[i][1]

		if intHi < newLo {
			result = append(result, intervals[i])
			i++
			continue
		}

		if intLo > newHi {
			result = append(result, []int{newLo, newHi})
			result = append(result, intervals[i:]...)
			break
		}

		lenInt := intHi - intLo
		lenNew := newHi - newLo
		max := int(math.Max(float64(intHi), float64(newHi)))
		min := int(math.Min(float64(intLo), float64(newLo)))
		if max-min <= lenNew+lenInt {
			newLo = min
			newHi = max
		}

		if i == len(intervals)-1 {
			result = append(result, []int{newLo, newHi})
			break
		}
		i++
	}

	return result
}
