func maximumProduct(nums []int) int {
    m1,m2,m3:=math.MinInt32,math.MinInt32,math.MinInt32
    a1,a2:=math.MaxInt32,math.MaxInt32
    for _,v:=range nums{
        if v>m1{
            m1,m2,m3=v,m1,m2
        }else if v>m2{
            m2,m3=v,m2
        }else if v>m3{
            m3=v
        }
        if v<a1{
            a1,a2=v,a1
        }else if v<a2{
            a2=v
        }
    }
    if m1*m2*m3>m1*a1*a2{
        return m1*m2*m3
    }else{
        return m1*a1*a2
    }
}
