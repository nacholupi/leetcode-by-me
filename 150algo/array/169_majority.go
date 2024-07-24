package array

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)

	c := len(nums) / 2

	if c == 0 {
		return nums[0]
	}

	if nums[c] == nums[c-1] {
		return nums[c-1]
	}
	return nums[c+1]
}
