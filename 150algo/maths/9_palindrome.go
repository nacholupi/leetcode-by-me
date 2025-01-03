package maths

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	original := x
	result := 0

	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	return result == original
}
