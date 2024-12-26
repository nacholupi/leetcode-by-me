package intervals

import (
	"fmt"
)

func summaryRanges(nums []int) (result []string) {
	if len(nums) == 0 {
		return []string{}
	}

	lo := nums[0]
	up := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != up+1 {
			result = append(result, rangeFmt(lo, up))
			lo = nums[i]
		}
		up = nums[i]
	}

	result = append(result, rangeFmt(lo, up))
	return result
}

func rangeFmt(lo int, up int) string {
	if lo == up {
		return fmt.Sprintf("%d", lo)
	} else {
		return fmt.Sprintf("%d->%d", lo, up)
	}
}
