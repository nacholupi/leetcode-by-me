package twopointers

func maxArea(height []int) int {

	max := 0
	i := 0
	j := len(height) - 1

	for i < j {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}
	return max
}
