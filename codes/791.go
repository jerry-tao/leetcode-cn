func customSortString(S string, T string) string {
    m:=make(map[rune]int)
    res:=""
    for _,v:=range T{
        m[v]++
    }
    for _,v:=range S{
        if count,ok:=m[v];ok{
            for i:=0;i<count;i++{
                res += string(v)
                m[v]--
            }
        }
    }
    for k,v:=range m{
        if v!=0{
            for i:=0;i<v;i++{
                res += string(k)
            }
        }
    }
    return res
}
