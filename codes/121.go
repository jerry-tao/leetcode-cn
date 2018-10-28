func maxProfit(prices []int) int {
    res:=0
    min:=math.MaxInt16
    for _,v:=range prices{
        if v<min{
            min=v
        }
        if v-min>res{
            res=v-min
        }
    }
    return res
}
