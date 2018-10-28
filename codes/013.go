func romanToInt(s string) int {
    res:=0
    r:=[]rune(s)
    m:=map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    for i,v:= range r{
        if i==0 || m[v]<=m[r[i-1]]{res+=m[r[i]]}else{res+=m[r[i]]-2*m[r[i-1]]}
    }
    return res
}
