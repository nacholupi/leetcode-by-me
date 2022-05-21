package powertwo

import "math"

func IsPowerOfTwo(n int) bool {
	_, frac := math.Modf(math.Log2(float64(n)))
	return frac == 0
}
