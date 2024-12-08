package array

func canJumpII(nums []int) int {
	smallestJumps := 0
	farthest := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == end {
			smallestJumps++
			end = farthest
		}
	}
	return smallestJumps
}
