func countPrimes(n int) int {
    if n<3{
        return 0
    }
    res := 1

    for i:=3;i<n;i+=2{
        if check(i){
            res+=1
        }
    }
    return res

}

func check(n int)bool{
    for i:=2;i*i<=n;i++{
        if n%i==0{
            return false
        }
    }
    return true
}
