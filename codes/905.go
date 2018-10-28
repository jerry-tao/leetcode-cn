func sortArrayByParity(A []int) []int {
    res:=make([]int,len(A))
    i:=0
    j:=len(A)-1
    for _,v:=range A{
        if v%2==0{
            res[i] = v
            i++
        }else{
            res[j]=v
            j--
        }
    }
    return res
}
