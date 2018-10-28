func findRadius(houses []int, heaters []int) int {
    mi:= 1000000000
    ma:=0

    for _,v:=range houses{
        for _,m:=range heaters{
            mi = min(mi, abs(m-v))
        }
        ma = max(ma,mi)
         mi= 1000000000
    }
    return ma
}

func min(a,b int)int{
    if a<b{return a}
    return b
}
func max(a,b int)int{
        if a>b{return a}
    return b
}
func abs(a int)int{
    if a<0{return -a}
    return a
}
