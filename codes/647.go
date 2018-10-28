func countSubstrings(s string) int {
    if len(s)==0{return 0}
    res:=0
    for i:=0;i<len(s);i++{
        sj,sk:=i,i
        for sj>=0 && sk<len(s) && s[sj]==s[sk]{
            sj--
            sk++
            res++
        }
        dj,dk:=i,i+1
        for dj>=0 && dk<len(s) && s[dj]==s[dk]{
            dj--
            dk++
            res++
        }
    }
    return res
}
