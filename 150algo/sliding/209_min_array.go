package sliding

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	dis := len(nums) + 1

	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			if right-left < dis {
				dis = right - left
			}
			sum -= nums[left]
			left++
		}
	}

	if dis == len(nums)+1 {
		return 0
	}
	return dis
}
