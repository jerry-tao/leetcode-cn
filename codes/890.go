func findAndReplacePattern(words []string, pattern string) []string {
    res:=[]string{}
    r1:=[]rune(pattern)
    for _,str:=range words{
        check:=true
        r2:=[]rune(str)
        m:=make(map[rune]rune)
        m2:=make(map[rune]rune)
        for i,c:=range r2{
            if q,ok:=m[c];ok && q!=r1[i]{
                check=false
                break
            }
            m[c] = r1[i]
        }
        for i,c:=range r1{
             if q,ok:=m2[c];ok && q!=r2[i]{
                check=false
                break
            }
            m2[c] = r2[i]
        }
        for k,v:=range m{
            if k != m2[v]{
                check=false
                break
            }
        }
        if check{
            res = append(res,str)
        }
    }
    return res
}
