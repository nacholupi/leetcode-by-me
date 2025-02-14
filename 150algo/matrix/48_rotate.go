package matrix

func rotate(matrix [][]int) {

	n := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < n-i; j++ {

			topLeft := matrix[i][j]
			topRight := matrix[j][n-i]
			downRight := matrix[n-i][n-j]
			downLeft := matrix[n-j][i]

			matrix[i][j] = downLeft
			matrix[j][n-i] = topLeft
			matrix[n-i][n-j] = topRight
			matrix[n-j][i] = downRight
		}
	}
}
