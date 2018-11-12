func compareVersion(version1 string, version2 string) int {
    s1:=strings.Split(version1,".")
    s2:=strings.Split(version2,".")
    i:=0
    k1,k2:=0,0
    max:=func(a,b int) int {if a>b{return a}
                            return b}
    for i<max(len(s1),len(s2)){
        if i>=len(s2){
            k2=0
        }else{
            k2,_=strconv.Atoi(s2[i])
        }
        if i>=len(s1){
            k1 = 0
        }else{
            k1,_=strconv.Atoi(s1[i])
        }
        if k1==k2 { i++;continue}
        if k1 > k2 { return 1 }
        if k1 < k2 {return -1 }
     }
    return 0
}
