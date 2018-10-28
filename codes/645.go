func findErrorNums(nums []int) []int {
    tmp := make([]int,len(nums))
    res := []int{}
    for _,v := range nums{
        tmp[v-1]++
    }
    for i,v := range tmp{
        if v==2{
            i++
            res = append([]int{i},res...)
        }
        if v==0{
            i++
            res = append(res,i)
        }
    }
    return res
}
