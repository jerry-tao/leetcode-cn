func optimalDivision(nums []int) string {
    res:=""
    num:=make([]string,len(nums))
    for i,v:=range nums{
        num[i] = strconv.Itoa(v)
    }
    res+=num[0]
    if len(num)>2{
    res+="/("+strings.Join(num[1:],"/")+")"
    }else if len(num)>1{
        res+="/"+num[1]
    }
    return res
}
