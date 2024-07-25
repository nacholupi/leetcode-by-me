package array

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	k = len(nums) - (k % (len(nums)))
	copy(nums, append(nums[k:], nums[:k]...))
}
