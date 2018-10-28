func detectCapitalUse(word string) bool {
    r:=[]rune(word)
    for i:=0;i<len(r);i++{
        if r[i]>='A' && r[i]<='Z'{
            if i>0 && !(r[i-1] >='A' && r[i-1]<='Z'){
                return false
            }
            if i<len(r)-1 &&i!=0 && !(r[i+1] >='A' && r[i+1]<='Z'){
                return false
            }
        }
    }
    return true
}
