func hammingDistance(x int, y int) int {
    res:=0
    tmp:=x^y
    for tmp!=0{
        res++
        tmp &= (tmp-1)
    }
    return res
}
