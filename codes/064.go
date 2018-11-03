func minPathSum(grid [][]int) int {
    min:=func(a,b int)int{if a<b{return a};return b}
    y:=len(grid)
    x:=len(grid[0])
    dp:=make([][]int, y)
    for i:=range dp{ dp[i] = make([]int,x) }
    dp[0][0] = grid[0][0]
    for i:=1; i<y;i++{ dp[i][0] = grid[i][0]+dp[i-1][0] }
    for i:=1; i<x;i++{ dp[0][i] = grid[0][i]+dp[0][i-1] }
    for i:=1;i<y;i++{
        for j:=1;j<x;j++{
            dp[i][j] = grid[i][j]+min(dp[i-1][j],dp[i][j-1])
        }
    }
    return dp[y-1][x-1]
}
