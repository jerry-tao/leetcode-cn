func findMin(nums []int) int {
    min:=nums[0]
    for i:=1;i<len(nums);i++{
        if nums[i]<nums[i-1]{
            min = nums[i]
            break
        }
    }
    return min
}
