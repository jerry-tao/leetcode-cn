func integerReplacement(n int) int {
    if n<=1{
        return 0
    }
    count:=0
    count = change(n,count)
    return count
}

func change(n , count int)int{
    if n<=1{
        return count
    }
    count++
    if n % 2==1{
        count = min(change(n+1,count),change(n-1,count))
    }else{
        count = change(n/2,count)
    }
    return count
}

func min(a,b int)int{
    if a<b{return a}
    return b
}
