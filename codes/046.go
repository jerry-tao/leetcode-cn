func permute(nums []int) [][]int {
    res:=[][]int{}
    out:=[]int{}
    visited:=make([]int,len(nums))
    dfs(nums,&res,out,visited,0)
    return res
}

func dfs(nums []int,res *[][]int, out []int ,visited []int, count int){

    if count==len(nums){
        o := make([]int,count)
        copy(o,out)
        *res = append((*res), o)
        return
    }
    for i,_:=range visited{
        if visited[i]==0{
            visited[i]=1
            out = append(out,nums[i])
            dfs(nums,res,out,visited,count+1)
            out = out[:len(out)-1]
            visited[i]=0
        }
    }

}
