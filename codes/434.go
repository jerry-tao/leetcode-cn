func countSegments(s string) int {
    count,str:=0,strings.Split(s," ")
    for _,v:=range str{
        if len(v)>0{
            count++
        }
    }
    return count
}
