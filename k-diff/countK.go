package kdiff

func CountKDifference(nums []int, k int) int {
	m := make(map[int]int)
	r := 0
	for _, val := range nums {
		m[val] += 1
	}

	for _, val := range nums {
		r += m[val+k]
	}
	return r
}
