func dominantIndex(nums []int) int {
    max,maxIndex,s2:=0,0,0
    for i,v:=range nums{
        if v>max{
            max,s2=v,max
            maxIndex = i
        }else if v>s2{
            s2=v
        }
    }
    if s2==0|| max/s2>=2{
        return maxIndex
    }
    return -1
}
