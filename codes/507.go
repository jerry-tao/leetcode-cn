func checkPerfectNumber(num int) bool {
    cur :=1
    for i:=2;i*i<=num;i++{
        if num%i==0{
            cur+=i
            cur+=num/i
            if i*i==num{
                cur-=i
            }
        }
    }
    return cur==num && num!=1
}
