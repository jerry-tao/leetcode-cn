func myAtoi(str string) int {
    res:=[]rune{}
    s:=0
    r:=0
    pos:=true
    var err error
    for _,c:=range str{
        if s==0{
            if c!=' '{
                if c=='-'{
                    pos=false
                    s=1
                }else if c>='0' && c<='9'{
                    res = append(res,c)
                    s=1
                }else if c=='+'{
                    s=1
                }else{
                    return 0
                }
            }
        }else if s==1 && c>='0'&&c<='9' {
                res = append(res,c)
        }else{
            break
        }
    }
    if len(res)>0{
        r,err=strconv.Atoi(string(res))
        if err!=nil || r>math.MaxInt32{
            if pos{r=math.MaxInt32}else{r=math.MinInt32}
        }else{
            if !pos{r=-r}
        }
    }
    return r
}
