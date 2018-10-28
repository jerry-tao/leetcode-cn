func findMaxConsecutiveOnes(nums []int) int {
    max:=0
    t:=0
    for _,v:=range nums{
        if v==1{
            t++
            if t>max{
                max=t
            }
        }else{
            t=0
        }
    }
    return max
}
