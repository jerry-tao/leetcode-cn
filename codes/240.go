func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0])==0{
        return false
    }
    x := len(matrix)-1
    y := 0
    for ;;{
        if matrix[x][y] == target{
            return true
        }else if matrix[x][y] > target{
            x--
        }else{
            y++
        }
        if x < 0 || y >=len(matrix[0]) {
            break
        }

    }
    return false
}
