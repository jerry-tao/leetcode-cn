func strStr(haystack string, needle string) int {
    if len(needle)==0{return 0}
    if len(haystack)==0{return -1}
    r1 :=[]rune(haystack)
    r2 :=[]rune(needle)
    for i,v:=range r1{
        if v==r2[0]{
            flag:=true
            for j,value:=range r2{
                if  i+j>=len(r1)||value!=r1[i+j]{
                    flag=false
                    break
                }
            }
            if flag{
                return i
            }
        }

    }
    return -1
}
