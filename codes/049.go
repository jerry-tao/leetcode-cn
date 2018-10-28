func groupAnagrams(strs []string) [][]string {
    res:=make([][]string,0)
    m:=make(map[string][]string)
    for _,s:=range strs{
        cnt:=make([]int,26)
        for _, c :=range s{
            cnt[int(c-'a')] += 1
        }
        key:=""
        for _,count:=range cnt{key+=strconv.Itoa(count)}
        if r,ok:=m[key];ok{
            m[key] = append(r,s)
        }else{
            m[key] = []string{s}
        }
    }
    for _,v:=range m{
        res = append(res,v)
    }
    return res

}
