func findRestaurant(list1 []string, list2 []string) []string {
    m:=make(map[string]int)
    var res []string
    min := 10000
    for i,v := range list1{
        m[v] = i
    }
    for i,v := range list2{
        if index,ok:=m[v];ok{
            if index+i==min{
                res = append(res,v)
            }
            if index+i<min{
                min = index+i
                res = []string{v}
            }
        }
    }
    return res
}
