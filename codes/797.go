func allPathsSourceTarget(graph [][]int) [][]int {
    res:=[][]int{}
    out:=[]int{0}
    n:=len(graph)
    _c := func(i []int)[]int{
        o:=make([]int,len(i))
        copy(o,i)
        return o
    }
    var dfs func(int)
    dfs = func(i int){
        if i==n-1{
            res = append(res,_c(out))
            return
        }
        j:=graph[i]
        for _,v:=range j{
            out = append(out, v)
            dfs(v)
            out = out[:len(out)-1]
        }
    }
    dfs(0)
    return res
}
