func backspaceCompare(S string, T string) bool {
    r1,r2 := []rune{},[]rune{}
    for _,v:=range S{
        if v=='#' {
            if len(r1)>0{r1 = r1[:len(r1)-1]}

        }else{
            r1 = append(r1, v)
        }
    }
    for _,v:=range T{
        if v=='#'{
            if  len(r2)>0 {r2 = r2[:len(r2)-1]}
        }else{
            r2 = append(r2, v)
        }
    }
    return string(r1)==string(r2)
}
