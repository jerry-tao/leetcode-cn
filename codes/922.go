func sortArrayByParityII(A []int) []int {
    i:=0
    j:=1
    res := make([]int,len(A))
    for _,v:=range A{
        if v%2==0{
            res[i] = v
            i+=2
        }else{
            res[j] = v
            j+=2
        }
    }
    return res
}
