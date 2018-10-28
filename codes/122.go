func maxProfit(prices []int) int {
    res:=0
    for i,_:=range prices{
        if i>0&&prices[i]>prices[i-1]{
            res+=prices[i]-prices[i-1]
        }
    }
    return res
}
