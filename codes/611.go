func triangleNumber(nums []int) int {
    count:=0
    for a:=0;a<len(nums);a++{
        for b:=a+1;b<len(nums);b++{
            for c:=b+1;c<len(nums);c++{
                if nums[a]+nums[b]>nums[c] && nums[b]+nums[c]>nums[a] && nums[a]+nums[c]>nums[b] && nums[a]!=0 && nums[b]!=0 && nums[c]!=0{
                    count+=1
                }
            }
        }
    }
    return count
}
