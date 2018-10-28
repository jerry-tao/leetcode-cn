func subsetsWithDup(nums []int) [][]int {
    res:=make([][]int,1)
    last:=nums[0]
    sort.Ints(nums)
    size:=1
    for i:=0;i<len(nums);i++{
        if last!=nums[i]{
            last=nums[i]
            size=len(res)
        }
        count:=len(res)
        for j:=count-size;j<count;j++{
            res=append(res,append([]int{nums[i]},res[j]...))
        }
    }
    return res
}
