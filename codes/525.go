func findMaxLength(nums []int) int {
    m:=make(map[int]int)
    var res int
    sum:=0
    m[0]=-1
    for i,v:=range nums{
        if v==1{
            sum+=1
        }else{
            sum-=1
        }
        if _,ok:=m[sum];ok{
            res = max(res,i-m[sum])
        }else{
            m[sum] = i
        }
        fmt.Println(m)
    }
    return res
}

func max(a,b int) int {
    if a>b{return a}
    return b
}
