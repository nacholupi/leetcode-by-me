package hashmap

import (
	"math"
)

func isHappy(n int) bool {

	set := make(map[int]struct{})
	in := n

	for {
		if _, ok := set[in]; ok {
			return false
		}

		set[in] = struct{}{}

		acum := 0
		for in != 0 {
			acum += int(math.Pow(float64(in%10), 2))
			in /= 10
		}

		in = acum

		if in == 1 {
			return true
		}
	}
}
