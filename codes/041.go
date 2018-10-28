func firstMissingPositive(nums []int) int {
    l:=len(nums)
    n:=make([]int,l+1)
    for _,v:=range nums{
        if v<=l&&v>0{
            n[v-1] = 1
        }
    }
    for i,v:=range n{
        if v==0{
            return i+1
        }
    }
    return 1
}
