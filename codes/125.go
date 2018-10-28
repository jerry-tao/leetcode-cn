func isPalindrome(s string) bool {
    r:=[]rune(s)
    left,right:=0,len(r)-1
    for left<right{
        for ;left<right;left++{
            if (r[left]>=97 && r[left]<=122) ||(r[left]>=65 && r[left]<=90) || (r[left]>=48 && r[left]<=57){
                break
            }
        }
        for ;left<right;right--{
            if (r[right]>=97 && r[right]<=122)|| (r[right]>=65 && r[right]<=90)|| (r[right]>=48 && r[right]<=57) {
                break
            }
        }
        if r[left]==r[right] || (r[left]-r[right]==32 && r[right]>=65) ||(r[left]-r[right]==-32 && r[left]>=65){
            left++
            right--
        }else{
            return false
        }
    }
    return true
}
