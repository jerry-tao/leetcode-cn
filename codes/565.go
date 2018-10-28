func arrayNesting(nums []int) int {
    res:=0
    for i:=0;i<len(nums);i++{
        count := 0
        j:=i
        for count==0 || j!=i{
            j = nums[j]
            count++
        }
        if count>res{
            res=count
        }
    }
    return res
}
