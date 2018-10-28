func moveZeroes(nums []int)  {
    s:=0
    f:=0
    for f<len(nums){

        if nums[f]!=0{
            nums[s],nums[f] = nums[f],nums[s]
            s++
        }
        f++


    }
}
