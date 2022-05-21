package array

func maxSubArray(nums []int) int {
	sum := 0
	grtst := 0
	for i := range nums {
		if i == 0 {
			sum = nums[i]
			grtst = nums[i]
			continue
		}

		sum = max(nums[i], sum+nums[i])

		if sum > grtst {
			grtst = sum
		}
	}
	return grtst
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
