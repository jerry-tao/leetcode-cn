func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }
    l:=0
    r:=len(nums)-1
    for l<r{
        mid := (l+r)/2
        if nums[mid]==target{return mid}
        if nums[mid]<nums[r]{
            if nums[r]>=target && nums[mid]<target{
                l=mid+1
            }else{
                r=mid-1
            }
        }else{
            if nums[l]<=target && nums[mid]>target{
                r=mid-1
            }else{
                l=mid+1
            }
        }
    }
    if nums[l]==target{
        return l
    }
    return -1
}
