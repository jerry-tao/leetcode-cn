func isPowerOfFour(num int) bool {
    cur:=1
    for cur<=num{
        if cur==num{return true}
        if cur>num{return false}
        cur*=4
    }
    return false
}
