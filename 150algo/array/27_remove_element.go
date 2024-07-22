package array

import "log"

func removeElement(nums []int, val int) int {

	count := 0
	tail := len(nums)
	for i := 0; i < tail; i++ {
		log.Println(nums)
		if nums[i] == val {
			tail--
			nums[i] = nums[tail]
			i--
			continue
		}
		count++

	}
	return count
}
