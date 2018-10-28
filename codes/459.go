func repeatedSubstringPattern(s string) bool {
    l := len(s)
    r := []rune(s)
    if l<2{
        return false
    }
    for i:=2;i<=l;i++{
        if l%i!=0{continue}
        tmp:=string(r[:l/i])
        c:=true
        for j:=l/i;j<=l-l/i;j+=l/i{
            if tmp!=string(r[j:j+l/i]){
                c=false
                break
            }
        }
        if c {
            return true
        }
    }
    return false
}
