func isValidSudoku(board [][]byte) bool {
    	_valid := func(i, j int) bool {
		for col := 0; col < 9; col++ {
			if col != j && board[i][j] == board[i][col] {
				return true
			}
		}
		for row := 0; row < 9; row++ {
			if row != i && board[i][j] == board[row][j] {
				return true
			}
		}
		for row := i / 3 * 3; row < i/3*3+3; row++ {
			for col := j / 3 * 3; col < j/3*3+3; col++ {
				if (row != i || col != j) && board[i][j] == board[row][col] {
					return true
				}
			}
		}
		return false;
	}
    for i:=0;i<9;i++{
        for j:=0;j<9;j++{
            if board[i][j]!='.' && _valid(i,j){
                return false
            }
        }
    }
    return true
}
