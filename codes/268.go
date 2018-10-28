func missingNumber(nums []int) int {
    result:=(0+len(nums))*(len(nums)+1)/2
    count:=0
    for _,v:=range nums{
        count+=v
    }
    return result-count
}
