func findUnsortedSubarray(nums []int) int {
    left,right:=0,len(nums)-1
    tmp := make([]int,len(nums))
    copy(tmp,nums)
    sort.Ints(tmp)
    for left<len(nums){
        if tmp[left] != nums[left]{
            break
        }
        left++
    }
    for right>=left{
        if tmp[right] != nums[right]{
            break
        }
        right--
    }
    return right-left+1
}
