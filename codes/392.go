func isSubsequence(s string, t string) bool {
    i:=0
    n:=len(s)
    if n==0{return true}
    for _,v:=range t{
        if v==rune(s[i]){
            i++
            if i==n{return true}
        }
    }
    return false
}
