func findShortestSubArray(nums []int) int {
    m:=make(map[int][]int)
    for i,v:=range nums{
        if _,ok:=m[v];ok{
            m[v][0]++
            m[v][2]=i
        }else{
            m[v] = []int{1,i,i}
        }
    }
    max:=0
    ml:=50000
    for _,v:=range m{
        if v[0]>max || (v[0]==max && v[2]-v[1]+1<ml) {
              max = v[0]
              ml = v[2]-v[1]+1
        }
    }
    return ml
}
