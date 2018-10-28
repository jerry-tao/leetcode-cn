func numberOfLines(widths []int, S string) []int {
    count :=0
    for _,v:=range S{
        if 100-count%100>=widths[int(v-'a')]{
            count+=widths[int(v-'a')]
        }else{
            count+=widths[int(v-'a')] + (100-count%100)
        }

    }
    fmt.Println(count)
    if count%100==0{
        return []int{count/100,0}
    }else{
    return []int{count/100+1,count%100}
    }

}
