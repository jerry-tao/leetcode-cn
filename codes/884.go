func uncommonFromSentences(A string, B string) []string {
    s1:=strings.Split(A," ")
    s2:=strings.Split(B," ")
    m:=make(map[string]int)
    res:=[]string{}
    for _,v:=range s1{
        m[v]+=1
    }
    for _,v:=range s2{
        m[v]+=1
    }
    for k,v:=range m{
        if v==1 && len(k)>0{
            res = append(res,k)
        }
    }
    return res
}
