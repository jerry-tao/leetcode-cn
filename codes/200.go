func numIslands(grid [][]byte) int {
    h:=len(grid)
    if h==0{return 0}
    w:=len(grid[0])
    visited:=make([][]byte,h)
    res:=0
    for i :=range visited{ visited[i] = make([]byte,w) }
    var dfs func(int,int)
    dfs=func(a,b int){
        if a>=h || a<0 || b>=w || b<0 || grid[a][b]==byte('0') || visited[a][b]==1{ return }
        visited[a][b]=1
        dfs(a,b+1)
        dfs(a,b-1)
        dfs(a+1,b)
        dfs(a-1,b)
    }
    for i,v:=range grid{
        for j,item:=range v{
            if item==byte('1') && visited[i][j]!=1{
                res+=1
                dfs(i,j)
            }
        }
    }
    return res
}
