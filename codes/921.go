func minAddToMakeValid(S string) int {
    lc,rc:=0,0
    for _,v:=range S{
        if v=='('{
            lc++
        }else{
            if lc>0{
                lc--
            }else{
                rc++
            }
        }
    }
    return lc+rc
}
