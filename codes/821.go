func shortestToChar(S string, C byte) []int {
    res:=make([]int,len(S))
    c:=-1
    for i:=range S{
        if S[i]==C{
            res[i] = 0
            c=i
        }else{
            if c==-1{
                res[i] = 10000
            }else{
                res[i] =  i-c
            }
        }
    }
    min:=func(a,b int) int{if a<b{return a};return b}
    c=-1
    for j:=len(S)-1;j>=0;j--{
        if S[j]==C{
            c=j
        }else{
            if c!=-1{
                res[j] = min(res[j], c-j)
            }
        }
    }

    return res
}
