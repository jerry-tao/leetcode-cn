func removeDuplicates(nums []int) int {
    s:=0
    f:=1
    count:=0
    for f<len(nums){
        if nums[f]==nums[s] && count==1{
            f++
        }else{
            if nums[f]==nums[s] {
                count++
            }else{
                count=0
            }
            nums[s+1] = nums[f]
            s++
            f++
        }
    }
    return s+1
}
