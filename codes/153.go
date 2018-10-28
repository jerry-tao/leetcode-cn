func findMin(nums []int) int {
    if len(nums) == 0{
        return -1
    }

    min:=nums[0]
    for _,v:=range nums{
        if v < min{
            return v
        }
    }
    return min
}
