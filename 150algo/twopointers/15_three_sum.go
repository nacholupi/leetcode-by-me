package twopointers

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	results := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		mp := make(map[int]int)

		for j := i + 1; j < len(nums); j++ {

			sum := nums[i] + nums[j]

			if v, ok := mp[nums[j]]; ok {
				results = append(results, []int{nums[i], v, nums[j]})
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			} else {
				// complement:= -sum
				mp[-sum] = nums[j]
			}

		}
	}
	return results
}
