func isAnagram(s string, t string) bool {
    a:=make([]int,26)
    b:=make([]int,26)

    for _,v:=range s{a[int(v-'a')]++}
    for _,v:=range t{b[int(v-'a')]++}
    k1,k2:="",""
    for _,v:=range a{k1+=strconv.Itoa(v)}
    for _,v:=range b{k2+=strconv.Itoa(v)}
    return k1==k2

}
