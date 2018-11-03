func minimumTotal(triangle [][]int) int {
    min:=func(a,b int) int{
        if a<b{return a}
        return b
    }

    dp:=make([][]int,len(triangle))
    for i:=range dp{ dp[i] = make([]int,len(triangle[i])) }
    dp[0][0] = triangle[0][0]
    for i:=1;i<len(triangle);i++{
        dp[i][0] = triangle[i][0]+dp[i-1][0]
        dp[i][len(triangle[i])-1] =  triangle[i][len(triangle[i])-1] + dp[i-1][len(triangle[i-1])-1]
    }
    for i:=1;i<len(triangle);i++{
        for j:=1;j<len(triangle[i])-1;j++{
            dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])
        }
    }
    res:=math.MaxInt32
    for _,v:=range dp[len(triangle)-1]{
        if v<res{
            res = v
        }
    }
    return res
}
