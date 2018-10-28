func titleToNumber(s string) int {
    res:=0
    l:=len(s)-1
    for _,v:=range s{
        res+=(int(v)-64)*pow(26, l)
        l--
    }
    return res
}
func pow(x, n int) int {
    ret := 1
    for n != 0 {
        if n%2 != 0 {
            ret = ret * x
        }
        n /= 2
        x = x * x
    }
    return ret
}
