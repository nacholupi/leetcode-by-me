package hashmap

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	maxResult := 0

	for _, n := range nums {
		if m[n] != 0 {
			continue
		}
		m[n] = m[n-1] + m[n+1] + 1
		if m[n-1] != 0 {
			m[n-1-(m[n-1]-1)] = m[n]
		}
		if m[n+1] != 0 {
			m[n+1+m[n+1]-1] = m[n]
		}

		if m[n] > maxResult {
			maxResult = m[n]
		}
	}
	return maxResult
}
