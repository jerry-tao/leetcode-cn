func findDuplicate(paths []string) [][]string {
    m:=make(map[string][]string)
    res := make([][]string,0)
    var tmp []string
    middle:=0
    for _,v:=range paths{
        tmp = strings.Split(v," ")
        for i:=1;i<len(tmp);i++{
            middle = strings.Index(tmp[i],"(")
            if _,ok:=m[tmp[i][middle:]];!ok{
                m[tmp[i][middle:]] = []string{}
            }
            m[tmp[i][middle:]] = append(m[tmp[i][middle:]], tmp[0]+"/"+tmp[i][:middle])
        }
    }
    for _,v:=range m{
        if len(v)>1{
        res = append(res,v)

        }
    }
    return res
}
