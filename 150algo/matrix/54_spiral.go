package matrix

func spiralOrder(matrix [][]int) []int {

	xi := 0
	xl := len(matrix)
	yi := 0
	yl := len(matrix[0])
	mxn := xl * yl
	xl--
	yl--

	result := make([]int, 0, mxn)
	direction := 0

	for len(result) < mxn {
		// left --> right
		if direction == 0 {
			for i := yi; i <= yl; i++ {
				result = append(result, matrix[xi][i])
			}
			xi++
		}

		// top -->down
		if direction == 1 {
			for i := xi; i <= xl; i++ {
				result = append(result, matrix[i][yl])
			}
			yl--
		}

		// right --> left
		if direction == 2 {
			for i := yl; i >= yi; i-- {
				result = append(result, matrix[xl][i])
			}
			xl--
		}

		// bottom --> up
		if direction == 3 {
			for i := xl; i >= xi; i-- {
				result = append(result, matrix[i][yi])
			}
			yi++
		}

		direction = (direction + 1) % 4
	}

	return result
}
