func binaryGap(N int) int {
    res:=0
    cur:=-1
    for i:=31;i>=0;i--{
        if N & (1<<uint(i)) != 0 {
            if cur==-1{ cur=i }else{
                if cur-i>res{res=cur-i}
                cur=i
            }
        }

    }
    return res
}
