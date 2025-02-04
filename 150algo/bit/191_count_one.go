package bit

func hammingWeight(n int) int {
	r := 0
	for n > 0 {
		r += n & 1
		n = n >> 1
	}
	return r
}
