func longestCommonPrefix(strs []string) string {
    if len(strs)==0{return ""}
    res:=""
    for i:=0;i<len(strs[0]);i++{
        c := strs[0][i]
        s:=false
        for _,v:=range strs{
            if i>=len(v)||v[i]!=c{
                s=true
                break
            }
        }
        if s{ return res }else{res+=string(c)}
    }
    return res
}
