func climbStairs(n int) int {
    if n<1{return 0}
    if n==1{return 1}
    if n==2{return 2}
    a:=1
    b:=2
    res:=2
    for i:=3;i<=n;i++{
        res=a+b
        a=b
        b=res
    }
    return res
}
