package maths

func mySqrt(x int) int {
	for i := 0; i < 50000; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return 0
}
