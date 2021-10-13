package divide

func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}

	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}
		if board[row][i] == c {
			return false
		}
		//block
		//行(row/3)*3+i/3
		//列(col/3)*3+i%3
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
