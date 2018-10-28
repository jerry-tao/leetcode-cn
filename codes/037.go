func solveSudoku(board [][]byte) {
	_valid := func(i, j int) bool {
		for col := 0; col < 9; col++ {
			if col != j && board[i][j] == board[i][col] {
				return false
			}
		}
		for row := 0; row < 9; row++ {
			if row != i && board[i][j] == board[row][j] {
				return false
			}
		}
		for row := i / 3 * 3; row < i/3*3+3; row++ {
			for col := j / 3 * 3; col < j/3*3+3; col++ {
				if (row != i || col != j) && board[i][j] == board[row][col] {
					return false
				}
			}
		}
		return true;
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == 9 {
			return true
		}
		if j >= 9 {
			return dfs(i+1, 0)
		}
		if board[i][j] == '.' {
			for k := 1; k <= 9; k++ {
				board[i][j] = byte(k) + '0'
				if _valid(i, j) {
					if dfs(i, j+1) {
						return true
					}
				}
				board[i][j] = '.'
			}
		} else {
			return dfs(i, j+1)
		}
		return false
	}
	dfs(0, 0)
}
