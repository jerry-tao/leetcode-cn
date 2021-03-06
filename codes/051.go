func solveNQueens(n int) [][]string {
    res:=[][]string{}
    	post := make([]int, n)
	for i := 0; i < n; i++ {
		post[i] = -1
	}
	dfs(post, 0, &res)
	return res
}

func dfs(pos []int, row int, res *[][]string) {
	n := len(pos)
	if row == n {
        out:=make([]string,n)
        for j,v:=range pos{
                a:=""
            for i:=0;i<n;i++{
                if i==v{
                    a+="Q"
                }else{
                    a+="."
                }
            }
            out[j] = a
        }
        *res  = append(*res,out)
	} else {
		for col := 0; col < n; col++ {
			if valid(pos, row, col) {
				(pos)[row] = col;
				dfs(pos, row+1, res)
				(pos)[row] = -1
			}
		}
	}
}
func valid(pos []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if col == pos[i] || abs(row-i) == abs(col-pos[i]) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
