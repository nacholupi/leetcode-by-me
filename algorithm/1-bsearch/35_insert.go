package bsearch

func searchInsert(nums []int, target int) int {
	subEnd := len(nums) - 1
	subInit := 0
	h := 0

	for {
		h = (subInit + subEnd) / 2
		if nums[h] == target {
			return h
		}

		if target > nums[h] {
			subInit = h + 1
		}

		if target < nums[h] {
			subEnd = h - 1
		}
		if subInit > subEnd {
			return subInit
		}
	}
}
