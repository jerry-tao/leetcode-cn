func hIndex(citations []int) int {
    sort.Ints(citations)
    n:=len(citations)
    for i,v:=range citations{
        if n-i<=v{
            return n-i
        }
    }
    return 0

}
