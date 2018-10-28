func findPeakElement(nums []int) int {
    res := 0
    max := nums[0]
    if len(nums)==1{
        return 0
    }
    for i:=1;i<len(nums);i++{
        if nums[i]>max{
            max = nums[i]
            res = i
        }else{
            res = i-1
            break
        }
    }
    return res
}
