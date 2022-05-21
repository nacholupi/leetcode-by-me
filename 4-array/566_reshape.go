package array

func matrixReshape(mat [][]int, r int, c int) [][]int {

	matR := len(mat)

	if matR == 0 {
		return [][]int{}
	}

	matC := len(mat[0])

	if matR*matC != r*c || r == matR && c == matC {
		return mat
	}

	res := make([][]int, r)
	for i := 0; i < matR*matC; i++ {
		if i%c == 0 {
			res[i/c] = make([]int, c)
		}
		res[i/c][i%c] = mat[i/matC][i%matC]
	}

	return res
}
