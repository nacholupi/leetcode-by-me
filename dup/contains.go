package dup

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		} else {
			m[val] = struct{}{}
		}
	}
	return false
}
