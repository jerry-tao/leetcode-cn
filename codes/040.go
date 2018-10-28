func combinationSum2(candidates []int, target int) [][]int {
        res:=[][]int{}
    out:=[]int{}
    var dfs func(int,int)
    sort.Ints(candidates)
    dfs=func(index int, sum int){
        if sum==0{
            res = append(res, c(out))
            return
        }
        for j:=index;j<len(candidates);j++{
            if sum-candidates[j]>=0 {
                if j>index&&candidates[j]==candidates[j-1]{continue}
                out = append(out,candidates[j])
                dfs(j+1,sum-candidates[j])
                out = out[:len(out)-1]
            }
        }
    }
    dfs(0,target)
    return res
}
func c(i []int)[]int{
    c:=make([]int, len(i))
    copy(c,i)
    return c
}
