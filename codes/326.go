func isPowerOfThree(n int) bool {
    cur:=1
    for cur<=n{
        if cur==n{return true}
        cur*=3
    }
    return false
}
