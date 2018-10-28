func repeatedStringMatch(A string, B string) int {
    i:=1
    la,lb:=len(A),len(B)
    if la>lb{
        if strings.Contains(A,B){
            return 1
        }
        if strings.Contains(A+A,B){
            return 2
        }
    }
    tmp := A
    for la<3*lb{
        if strings.Contains(tmp,B){
            return i
        }
        tmp+=A
        la = len(tmp)
        i++
    }
    return -1
}
