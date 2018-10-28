func findMaxAverage(nums []int, k int) float64 {
    max:=0
    for i:=0;i<k;i++{
        max += nums[i]
    }
    cur:=max
    for i:=k;i<len(nums);i++{
        tmp:=cur+nums[i]-nums[i-k]
        if tmp>max{
            max=tmp
        }
        cur=tmp
    }
    return float64(max)/float64(k)

}
