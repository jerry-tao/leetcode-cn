func findLengthOfLCIS(nums []int) int {
    if len(nums)==0{
        return 0
    }
    j:=0
    max:=1
    for i:=1;i<len(nums);i++{
        if nums[i]>nums[i-1]{
            if i-j+1>max{
                max=i-j+1
            }
        }else{
            j=i
        }
    }
    return max
}
