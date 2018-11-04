func canJump(nums []int) bool {
    reach:=0
    for i:=0;i<len(nums);i++{
        if reach>=len(nums)-1 || reach < i{break}
        if i+nums[i]>reach{
            reach = i+nums[i]
        }
    }
    return reach>=len(nums)-1
}
