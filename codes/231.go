func isPowerOfTwo(n int) bool {
    cur :=1
    for cur<=n{
        if cur==n{
            return true
        }else if cur>n{
            return false
        }else{
        cur=cur*2
        }

    }
    return false
}
