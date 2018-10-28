func partition(s string) [][]string {
    res := [][]string{}
    out := []string{}
    var dfs func(int)
    dfs = func(start int){
        if start==len(s){
            res = append(res, c(out))
            return
        }
        for i:=start;i<len(s);i++{
            if valid(s,start,i){
                out = append(out, s[start:i+1])
                dfs(i+1)
                out = out[:len(out)-1]
            }
        }
    }
    dfs(0)
    return res

}

func c(i []string)[]string{
    o := make([]string,len(i))
    copy(o,i)
    return o
}

func valid(s string, start int, end int) bool{
    for start < end {
        if s[start] != s[end]{ return false}
        start++
        end--
    }
    return true;
}
