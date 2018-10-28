func search(nums []int, target int) int {
    if len(nums)==1 && nums[0]==target{
        return 0
    }
    left:=0
    right:=len(nums)-1
    for left<right{
        mid:=(left+right)/2
        if nums[mid]>target{right=mid}
        if nums[mid]<target{left=mid+1}
        if nums[mid]==target{
            return mid
        }
    }
    if nums[left]==target{
        return left
    }
    return  -1
}
