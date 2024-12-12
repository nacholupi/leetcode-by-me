package array

func productExceptSelf(nums []int) []int {

	prefixProd := make([]int, len(nums))
	suffixProd := make([]int, len(nums))
	prefixProd[0] = nums[0]
	suffixProd[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		prefixProd[i] = prefixProd[i-1] * nums[i]

		j := len(nums) - 1 - i
		suffixProd[j] = suffixProd[j+1] * nums[j]
	}

	result := make([]int, len(nums))
	result[0] = suffixProd[1]
	result[len(nums)-1] = prefixProd[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		result[i] = prefixProd[i-1] * suffixProd[i+1]
	}

	return result
}
