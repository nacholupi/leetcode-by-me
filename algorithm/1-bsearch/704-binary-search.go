package bsearch

func search(nums []int, target int) int {
	l := len(nums)
	ini := 0
	end := l - 1

	for end-ini >= 0 {
		hl := (ini + end) / 2
		mVal := nums[hl]
		if mVal == target {
			return hl
		}
		if mVal < target {
			ini = hl + 1
		}
		if mVal > target {
			end = hl - 1
		}
	}

	return -1
}
