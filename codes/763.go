func partitionLabels(S string) []int {
    res:=[]int{}
    m:=make(map[rune]int)
    for i,v:=range S{m[v]=i}
    s,l:=0,0
    for i,v:=range S{
        if l<m[v]{l=m[v]}
        if i==l{
            res = append(res,i-s+1)
            s=i+1
        }
    }
    return res
}
