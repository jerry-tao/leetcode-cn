func longestPalindrome(s string) int {
    m:=make(map[rune]int)
    for _,v:=range s{
        m[v] ++
    }
    res:=0
    flag:=false
    for _,v := range m{
        if v!=0 && v%2==0{
            res+=v
        }else{
            res+=v-1
            flag=true
        }
    }
    if flag{res++}
    return res
}
