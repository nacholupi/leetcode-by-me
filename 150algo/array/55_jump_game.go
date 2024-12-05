package array

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

func canJumpRecursive(nums []int) bool {
	var canJumpFromPosition func(idx int) bool
	canJumpFromPosition = func(idx int) bool {
		if idx >= len(nums) {
			return false
		}

		if idx == len(nums)-1 {
			return true
		}

		if nums[idx] == 0 {
			return false
		}

		for i := 1; i <= nums[idx]; i++ {
			if canJumpFromPosition(idx + i) {
				return true
			}
		}
		return false
	}

	return canJumpFromPosition(0)
}
