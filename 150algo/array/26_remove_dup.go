package array

func removeDuplicates(nums []int) int {
	k := 1

	if len(nums) == 1 {
		return k
	}

	for i := 1; i < len(nums); i++ {
		if nums[k-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
