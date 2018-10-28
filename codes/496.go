func nextGreaterElement(findNums []int, nums []int) []int {
    res:=make([]int,len(findNums))
    for i,v:=range findNums{
        check:=0
        for _,value:=range nums{
            if v!=value && check==0{continue}
            if v==value{ check=1 }
            if value>v{ check=2;res[i]=value;break }
        }
        if check==1{res[i]=-1}
    }
    return res
}
