func convertToTitle(n int) string {
    res:=[]rune{}
    for n!=0{
        n--
        res = append([]rune{rune(n%26)+'A'},res...)
        n = n/26
    }
    return string(res)
}
