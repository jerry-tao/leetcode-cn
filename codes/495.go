func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 0{
        return 0
    }
    res:=duration
    for i:=1;i<len(timeSeries);i++{
        res+=duration
        if duration-(timeSeries[i]-timeSeries[i-1])>0{
          res-=duration-(timeSeries[i]-timeSeries[i-1])
        }

    }
    return res
}
