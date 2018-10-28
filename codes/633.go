func judgeSquareSum(c int) bool {
    m:=make(map[int]int)
    for i:=0;i*i<=c;i++{
        if _,ok:=m[i*i];ok||i*i*2==c{
            return true
        }else{
            m[c-i*i]=1
        }
    }
    return false
}
