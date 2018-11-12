func countPrimeSetBits(L int, R int) int {
    c :=0
    m:=map[int]bool{
        2: true,
        3: true,
        5: true,
        7: true,
        11: true,
        13: true,
        17: true,
        19: true,
    }
    count:=func(n int) int{
        res:=0
        for n!=0{
            res++
            n = n&(n-1)
        }
        return res
    }

    for i:=L;i<=R;i++{
        if m[count(i)] {
            c++
        }
    }
    return c
}
