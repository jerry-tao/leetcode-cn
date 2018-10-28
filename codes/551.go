func checkRecord(s string) bool {
    ca:=0
    cl:=0
    for _,r:=range s{
        if r=='A'{
            ca++
            if ca>1{return false}
        }
        if r=='L'{
            cl++
            if cl>2{return false}
        }else{
            cl=0
        }
    }
    return true
}
