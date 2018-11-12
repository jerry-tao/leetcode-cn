func getSum(a int, b int) int {
    if b==0{return a }
    s:= a^b
    c:= (a&b)<<1
    return getSum(s,c)
}
