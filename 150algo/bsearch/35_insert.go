package bsearch

func searchInsert(nums []int, target int) int {
	subEnd := len(nums) - 1
	subInit := 0

	for subInit <= subEnd {
		m := (subInit + subEnd) / 2
		if nums[m] == target {
			return m
		}

		if target > nums[m] {
			subInit = m + 1
		}

		if target < nums[m] {
			subEnd = m - 1
		}
	}
	return subInit
}
