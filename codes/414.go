func thirdMax(nums []int) int {
    a,b,c:= math.MinInt64,math.MinInt64,math.MinInt64 
    for _,v:=range nums{
        if v>c{
            if v>b{
                if v>a{
                    a,b,c=v,a,b
                }
                if v < a{
                    b,c = v,b
                }
            }
            if v<b{
                c=v
            }
        }
    }
    if c==math.MinInt64{
        return a
    }else{
        return c
    }
}
