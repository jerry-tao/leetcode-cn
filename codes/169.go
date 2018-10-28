func majorityElement(nums []int) int {
    m := make(map[int]int)
    for _,v := range nums{
        m[v]+=1
        if m[v]>len(nums)/2{
            return v
        }
    }
    return -1
}
