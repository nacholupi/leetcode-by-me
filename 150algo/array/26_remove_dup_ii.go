package array

func removeDuplicatesII(nums []int) int {
	k := 1

	if len(nums) == 1 {
		return k
	}

	count := 0

	for i := 1; i < len(nums); i++ {
		if nums[k-1] != nums[i] {
			count = 0
		} else {
			count++
		}

		if count < 2 {
			nums[k] = nums[i]
			k++
		}

	}
	return k
}
