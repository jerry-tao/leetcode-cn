func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    res:=0
    for k:=0;k<len(nums);k+=2{
        res += nums[k]
    }
    return res
}
