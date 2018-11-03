func numberOfArithmeticSlices(A []int) int {
    res:=0
    c:=0
    for i:=2;i<len(A);i++{
        if A[i] - A[i-1] == A[i-1] -A[i-2]{
            c++
            res+=c
        }else{
            c=0
        }
    }
    return res
}
