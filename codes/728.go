func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for ;left<=right;left++{
        if check(left){
            res=append(res,left)
        }
    }
    return res
}
func check(n int) bool{
    str:=strconv.Itoa(n)
    for _,c := range str{
        p,_:=strconv.Atoi(string(c))
        if p==0 ||n%p!=0{
            return false
        }
    }
    return true
}
