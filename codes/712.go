func minimumDeleteSum(s1 string, s2 string) int {
    min:=func(a,b int)int{
        if a<b{return a}
        return b
    }
    l1,l2 := len(s1), len(s2)
    dp:=make([][]int,l1+1)
    for i:=range dp{
        dp[i] = make([]int,l2+1)
    }
    for i:=1;i<=l2;i++{
        dp[0][i] = dp[0][i-1] + int(s2[i-1])
    }
    for i:=1;i<=l1;i++{
        dp[i][0] = dp[i-1][0] + int(s1[i-1])
    }
    for i:=1;i<=l1;i++{
        for j:=1;j<=l2;j++{
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = min(dp[i-1][j]+int(s1[i-1]),dp[i][j-1]+int(s2[j-1]))
            }
        }
    }

    return int(dp[l1][l2])

}
