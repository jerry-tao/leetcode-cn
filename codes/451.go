func frequencySort(s string) string {
    re:=make([]string,len(s)+1)
    m:=make(map[rune]int)
    res:=""
    for _,v:=range s{m[v]++}
    for k,v:=range m{
        re[v] += strings.Repeat(string(k),v)
    }
    for j:=len(s);j>=0;j--{
        res+=re[j]
    }
    return res
}
