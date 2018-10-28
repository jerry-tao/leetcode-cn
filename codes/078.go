func subsets(nums []int) [][]int {
    res:=make([][]int,1)
    for i:=0;i<len(nums);i++{
        s:=len(res)
        for j:=0;j<s;j++{
            tmp:=make([]int,len(res[j]))
            copy(tmp,res[j])
            res = append(res,tmp)
            res[len(res)-1] = append(res[len(res)-1],nums[i])
        }
    }
    return res
}
