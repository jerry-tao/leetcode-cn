func isToeplitzMatrix(matrix [][]int) bool {
    res := true
    for j:=len(matrix)-1;j>=0;j--{
        a:=j
        b:=0

        for a<len(matrix) && b<len(matrix[0]){
            if matrix[a][b]!=matrix[j][0]{
                return false
            }
            a++
            b++
        }
    }
    for i:=len(matrix[0])-1;i>0;i--{
        a:=i
        b:=0
        for a<len(matrix[0]) && b<len(matrix){
            if matrix[b][a]!=matrix[0][i]{
                return false
            }
            a++
            b++
        }
    }
    return res
}
