package array

func candy(ratings []int) int {
	sum := len(ratings)

	if sum <= 1 {
		return sum
	}

	extraR := make([]int, len(ratings))
	extraL := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {

		if i != 0 && ratings[i] > ratings[i-1] {
			extraR[i] += extraR[i-1] + 1
		}

		j := len(ratings) - 1 - i
		if j != len(ratings)-1 && ratings[j] > ratings[j+1] {
			extraL[j] += extraL[j+1] + 1
		}
	}

	for i := 0; i < len(ratings); i++ {
		sum += max(extraR[i], extraL[i])
	}

	return sum
}
