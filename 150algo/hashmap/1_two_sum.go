package hashmap

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := m[nums[i]]
		if ok {
			return []int{idx, i}
		}
		m[target-nums[i]] = i
	}

	return []int{}
}
