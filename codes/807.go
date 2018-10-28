func maxIncreaseKeepingSkyline(grid [][]int) int {
    res:=0
    maxRow:=make([]int,len(grid[0]))
    maxCol:=make([]int,len(grid))
    for a,row:=range grid{
        for b,col:=range row{
            if col>maxRow[a]{
                maxRow[a] = col
            }
            if col>maxCol[b] {
                maxCol[b] = col
            }
        }
    }
    for a,row:=range grid{
        for b,col:=range row{
            if maxRow[a]<maxCol[b]{
                res+=maxRow[a]-col
            }else{
                res+=maxCol[b]-col
            }
        }
    }
    return res
}
