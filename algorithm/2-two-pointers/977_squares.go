package twopointers

func sortedSquares(nums []int) []int {
	r := make([]int, len(nums))
	i, j := 0, len(nums)-1

	for k := len(nums) - 1; k >= 0; k-- {
		if sqrt(nums[i]) > sqrt(nums[j]) {
			r[k] = sqrt(nums[i])
			i++
		} else {
			r[k] = sqrt(nums[j])
			j--
		}
	}
	return r
}

func sqrt(i int) int {
	return i * i
}
