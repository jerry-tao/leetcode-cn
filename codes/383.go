func canConstruct(ransomNote string, magazine string) bool {
    m:=make(map[rune]int)
    for _,v:=range magazine{
        m[v]++
    }
    for _,v:=range ransomNote{
        if c,ok:=m[v];ok&&c!=0{
            m[v]--
        }else{
            return false
        }
    }
    return true
}
