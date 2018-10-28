func findTheDifference(s string, t string) byte {
    m:=make(map[byte]int)
    sb:=[]byte(s)
    tb:=[]byte(t)
    for _,v:=range sb{
        m[v]+=1
    }
    for _,v:=range tb{
        if m[v]>0{
            m[v]--
        }else{
            return v
        }
    }
    return 1
}
