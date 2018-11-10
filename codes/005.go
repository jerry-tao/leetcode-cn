func longestPalindrome(s string) string {
    l := len(s)
    if l==1 || l==0{
        return s
    }
    dp:=make([][]bool,l)
    for i:=range dp{
        dp[i] = make([]bool,l)
        dp[i][i] = true
    }
    max,left,right:=0,0,0
    for i:=0;i<l;i++{
        for j:=0;j<i;j++{
            dp[j][i] = (s[i] == s[j] && (i - j < 2 || dp[j + 1][i - 1]))
                if (dp[j][i] && max < i - j + 1) {
                    max = i - j + 1
                    left = j
                    right = i
                }
        }
    }
    return s[left:right+1]
}
