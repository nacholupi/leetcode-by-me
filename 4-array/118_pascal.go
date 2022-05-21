package array

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		if i == 0 {
			res[i] = []int{1}
		} else {
			res[i] = make([]int, i+1)
			for j := 0; j < len(res[i]); j++ {
				if j == 0 || j == len(res[i])-1 {
					res[i][j] = 1
				} else {
					res[i][j] = res[i-1][j-1] + res[i-1][j]
				}
			}
		}
	}

	return res
}
