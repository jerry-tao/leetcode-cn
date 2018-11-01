func integerBreak(n int) int {
    if n == 2 || n == 3 {return n-1}
    if n == 4 { return 4 }
    switch n%3{
        case 0: return int(math.Pow(3,float64(n/3)))
        case 1: return int(math.Pow(3,float64(n/3-1)))*4
        case 2: return int(math.Pow(3,float64(n/3)))*2
    }
    return 0
}
