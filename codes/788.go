func rotatedDigits(N int) int {
    res:=0
    for n:=1;n<=N;n++{
        s := strconv.Itoa(n)
        flag:=false
        for _,v:=range s{
            if v=='3'||v=='4'||v=='7'{flag=false;break}
            if v=='2'||v=='5'||v=='6'||v=='9'{flag=true}
        }
        if flag{res++}
    }
    return res
}
