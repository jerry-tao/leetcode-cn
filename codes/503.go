func nextGreaterElements(nums []int) []int {
    n:=len(nums)
    res:=make([]int,n)
    for i:=range res{res[i]=-1}
    s:=[]int{}
    for i:=0;i<2*n;i++{
        num:=nums[i%n]
        for len(s)!=0 && nums[s[len(s)-1]]<num{
            res[s[len(s)-1]] = num
            s = s[:len(s)-1]
        }
        if i < n {
            s = append(s, i)
        }
    }
    return res
}
