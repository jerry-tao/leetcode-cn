func singleNumber(nums []int) []int {
    m:=make(map[int]int)
    res:=[]int{}
    for _,v:=range nums{
        m[v]++
    }
    for k,v:=range m{
        if v==1{ res = append(res,k) }

    }
    return res
}
