func reverseWords(s string) string {
    str := strings.Split(s," ")
    // reverse(str)
    for i:=0;i<len(str);i++{
        bs:=[]byte(str[i])
        ts:=[]byte{}
        for j:=len(bs)-1;j>=0;j--{
            ts=append(ts,bs[j])
        }
        str[i] = string(ts)
    }
    s=strings.Join(str," ")
    return s
}
