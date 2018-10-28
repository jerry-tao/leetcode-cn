func largestNumber(nums []int) string {
    res:=make([]string,len(nums))
    for i,v:=range nums{
        res[i] = strconv.Itoa(v)
    }
    sort.Sort(StringSlice(res))
    r:=strings.Join(res,"")
    if r[0]=='0'{
        return "0"
    }
    return r
}

type StringSlice []string

func (s StringSlice)Len()int{
    return len(s)
}
func (s StringSlice)Less(a,b int) bool{
    return s[a]+s[b]>s[b]+s[a]
}

func (s StringSlice)Swap(a,b int){
    s[a],s[b] = s[b],s[a]
}
