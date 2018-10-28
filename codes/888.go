func fairCandySwap(A []int, B []int) []int {
    s1,s2:=0,0
    for _,v:=range A{s1+=v}
    for _,v:=range B{s2+=v}
    s1 = (s1-s2)/2
    m:=make(map[int]int)
    for _,v:=range B{m[v]=v}
    for _,v:=range A{
        if a,ok:=m[v-s1];ok{
            return []int{v,a}
        }
    }
    return nil
}
