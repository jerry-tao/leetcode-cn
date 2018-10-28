func max(a,b int)int{
    if a>b{return a}
    return b
}
func min(a,b int)int{
    if a<b{return a}
    return b
}
func smallestRangeII(A []int, K int) int {
    sort.Ints(A)
    var ma,mi,res int
    res = A[len(A)-1]-A[0]
    for index,value:= range A{
        if index+1>=len(A){
            return res
        }
        mi = min(A[0]+K,A[index+1]-K)
        ma = max(value+K,A[len(A)-1]-K)
        res = min(res, ma-mi)
    }
    return res
}
