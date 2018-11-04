func rand10() int {
    key:=100
    for key>40{
        key = 7*(rand7()-1)+rand7()
    }
    return key%10+1
}
