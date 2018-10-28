func dailyTemperatures(temperatures []int) []int {
    res:=make([]int,len(temperatures))
    for i,v:=range temperatures{
        for j:=i;j<len(temperatures);j++{
            if temperatures[j]>v{
                res[i]=j-i
                break
            }
        }
    }
    return res
}
