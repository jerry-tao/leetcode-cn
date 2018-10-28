func pivotIndex(nums []int) int {
    res := -1
    sum:=0
    csum:=0
    for _,v:=range nums{
        sum+=v
    }
    for i,v:=range nums{
        if (sum-v)/2==csum && (sum-v)%2==0{
            res=i
            break
        }
        csum+=v

    }
    return res
}
