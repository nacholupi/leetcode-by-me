package array

func romanToInt(s string) int {

	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	lastValue := 1
	sum := 0

	for i := 0; i < len(s); i++ {
		currentValue := symbols[s[len(s)-1-i]]

		if currentValue >= lastValue {
			sum += currentValue
		}

		if currentValue > lastValue {
			lastValue = currentValue
		}

		if currentValue < lastValue {
			sum -= currentValue
		}
	}

	return sum
}
