func productExceptSelf(nums []int) []int {
    n:=len(nums)
    res:=make([]int,n)
    cal:= 1
    res[0] = 1
    for i:=1;i<n;i++{
        cal = cal*nums[i-1]
        res[i] = cal
    }
    cal=1
    for i:=len(res)-2;i>=0;i--{
        cal = cal * nums[i+1]
        res[i] = cal*res[i]
    }
    return res
}
