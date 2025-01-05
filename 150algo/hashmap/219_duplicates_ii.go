package hashmap

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for j, n := range nums {
		i, ok := m[n]
		if ok && j-i <= k {
			return true
		}
		m[n] = j
	}
	return false
}
