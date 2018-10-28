func reverseVowels(s string) string {
    r:=[]rune(s)
    left,right:=0,len(r)-1
    for left<right{
        for ;left<right;left++{
            if strings.Contains("aeiouAEIOU",string(r[left])){
                break
            }
        }
        for ;left<right;right--{
            if strings.Contains("aeiouAEIOU",string(r[right])){
                break
            }
        }
        tmp:=r[left]
        r[left]=r[right]
        r[right] = tmp
        left++
        right--
    }
    return string(r)
}
