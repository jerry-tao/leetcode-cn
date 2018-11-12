func findComplement(num int) int {
    flag:=false
    for i:=31;i>=0;i--{
        if (num & (1<<uint(i)))!=0 {flag = true}
        if flag {num ^= (1<<uint(i))}
    }
    return num
}
