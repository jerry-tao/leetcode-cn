func findDisappearedNumbers(nums []int) []int {
    res:=[]int{}
    for i:=0;i<len(nums);i++{
        v:=nums[i]
        if nums[v-1]!=v{
        nums[i]=nums[v-1]
        nums[v-1] = v
            i--
            }
    }
    for i,v:=range nums{
        if i+1!=v{
            res = append(res,i+1)
        }
    }
    return res
}
