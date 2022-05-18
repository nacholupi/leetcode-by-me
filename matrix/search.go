package matrix

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	if r == 0 {
		return false
	}
	c := len(matrix[0])

	iSeg := 0
	eSeg := r*c - 1

	for eSeg-iSeg >= 0 {
		i := (eSeg-iSeg)/2 + iSeg

		mVal := matrix[i/c][i%c]

		if mVal == target {
			return true
		}

		if mVal > target {
			eSeg = i - 1
		} else {
			iSeg = i + 1

		}
	}

	return false
}
