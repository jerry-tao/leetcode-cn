func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for i,v:=range nums{
        if index,ok:=m[v];ok{
            return []int{index,i}
        }else{
            m[target-v]=i
        }
    }
    return nil
}
