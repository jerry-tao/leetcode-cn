func findMinDifference(timePoints []string) int {
    m:=make([]int,1440)
    for _,v:=range timePoints{
        s:=strings.Split(v,":")
        a,_:=strconv.Atoi(s[0])
        b,_:=strconv.Atoi(s[1])
        m[a*60+b]++
        if m[a*60+b]>1{
            return 0
        }
    }
    a,b:=0,0
    for _,v:=range m{
        if v!=0{
            break
        }
        a++
    }
    for j:=1439;j>=0;j--{
        b++
        if m[j]!=0{
            break
        }
    }
    min:=a+b
    c:=-1
    for i,v:=range m{
        if v==1 {
            if c==-1{
                c=i
            }else{
                if i!=c && i-c<min{
                    min=i-c
                }
                c=i
            }
        }
    }
    return min
}
