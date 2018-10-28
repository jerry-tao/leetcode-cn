func removeDuplicates(nums []int) int {
    s:=0
    f:=1
    for f<len(nums){
        if nums[s]==nums[f]{
            f++
        }else{
            if f-s>1 {
                nums[s+1]=nums[f]
            }
            f++
            s++
        }

    }
    return s+1
}
