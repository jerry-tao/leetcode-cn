func searchMatrix(matrix [][]int, target int) bool {
    for _,v:=range matrix{
            for _,q:= range v{
                if q==target{return true}
            }
    }
    return false
}
