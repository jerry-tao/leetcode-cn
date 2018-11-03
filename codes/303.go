type NumArray struct {
    dp []int
}


func Constructor(nums []int) NumArray {
    dp := make([]int,len(nums))
    copy(dp,nums)
    for i:=1;i<len(dp);i++{
        dp[i]+=dp[i-1]
    }
    return NumArray{dp: dp}
}


func (this *NumArray) SumRange(i int, j int) int {
    if i==0{
        return this.dp[j]
    }
    return this.dp[j] - this.dp[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
