func transpose(A [][]int) [][]int {
    res:= make([][]int,len(A[0]))
    for i:=0;i<len(res);i++{
        res[i] = make([]int,len(A))
    }
    for a:=0;a<len(A);a++{
        for b:=0;b<len(A[a]);b++{
            res[b][a] = A[a][b]
        }
    }
    return res
}
