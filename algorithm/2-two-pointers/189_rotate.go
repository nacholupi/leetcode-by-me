package twopointers

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	left := make([]int, len(nums)-k)
	copy(left, nums[:len(nums)-k])
	right := nums[len(nums)-k:]

	copy(nums, right)
	copy(nums[k:], left)
}
