func firstUniqChar(s string) int {
    r:=[]rune(s)
    m:=make(map[rune]int)
    for _,v:=range r{m[v]++}
    for i,v:=range r{
        if m[v]==1{
            return i
        }
    }
    return -1
}
