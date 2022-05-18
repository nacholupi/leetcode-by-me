package soduku

func isValidSudoku(board [][]byte) bool {
	return validRows(board) && validCols(board) && validSquares(board)
}

func validRows(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		if ok := validSection(board, r, 0, 9); !ok {
			return false
		}
	}
	return true
}

func validCols(board [][]byte) bool {
	for c := 0; c < 9; c++ {
		if ok := validSection(board, 0, c, 1); !ok {
			return false
		}
	}
	return true
}

func validSquares(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		if ok := validSection(board, r/3*3, r%3*3, 3); !ok {
			return false
		}
	}
	return true
}

func validSection(board [][]byte, initr, initc, colSection int) bool {
	m := make(map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		r := initr + i/colSection
		c := initc + i%colSection
		cellValue := board[r][c]
		if cellValue != '.' {
			if _, ok := m[cellValue]; ok {
				return false
			}
			m[cellValue] = struct{}{}
		}
	}
	return true
}
