func reverseStr(s string, k int) string {
    r:=[]rune(s)
    n := len(r)
    for i:=0;i<n;i+=k*2{
        left:= i
        right:=n-1
        if i+k<n{
            right=i+k-1
        }
        for left<right{
            r[left],r[right] = r[right],r[left]
            left++
            right--
        }
    }
    return string(r)
}
