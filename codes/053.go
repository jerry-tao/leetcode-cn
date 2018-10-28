func maxSubArray(nums []int) int {
    s:=0
    res:=math.MinInt32
    for _,v:=range nums{
        if s+v>v{
            s=s+v
        }else{
            s=v
        }
        if s>res{
            res=s
        }
    }
    return res
}
