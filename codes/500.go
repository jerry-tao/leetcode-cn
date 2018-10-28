func findWords(words []string) []string {
    res := []string{}
    for _,v:=range words{
        a,b,c:=0,0,0
        for _,i:=range v{
            if strings.Contains("qwertyuiopQWERTYUIOP",string(i)){a=1}
            if strings.Contains("asdfghjklASDFGHJKL",string(i)){b=1}
            if strings.Contains("zxcvbnmZXCVBNM",string(i)){c=1}
            if a+b+c>1{break}
        }
        if a+b+c==1{res=append(res,v)}
    }
    return res
}
