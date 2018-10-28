func flipAndInvertImage(A [][]int) [][]int {
    res := make([][]int, len(A))
    for i,_:=range A{
        for j:=len(A[i])-1;j>=0;j--{
            if A[i][j] == 0{
              res[i] = append(res[i],1)
            }else{
                res[i] = append(res[i],0)
            }

        }
    }
    return res
}
