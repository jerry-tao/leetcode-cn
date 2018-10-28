func imageSmoother(M [][]int) [][]int {
    res:=make([][]int,len(M))
    for i:=0;i<len(res);i++{
        res[i]=make([]int,len(M[0]))
    }
    for a,sub:=range M{
        for b,v:=range sub{
            count:=1
            sum:=v
            if a-1>=0{
                count++
                sum+=M[a-1][b]
            }
            if a+1<len(M){
                count++
                sum+=M[a+1][b]
            }
            if b-1>=0 && a-1>=0{
                count++
                sum+=M[a-1][b-1]
            }
            if b-1>=0 {
                count++
                sum+=M[a][b-1]
            }
            if b-1>=0 && a+1<len(M) {
                count++
                sum+=M[a+1][b-1]
            }
            if b+1<len(sub) && a-1>=0{
                count++
                sum+=M[a-1][b+1]
            }
            if b+1<len(sub) {
                count++
                sum+=M[a][b+1]
            }
            if b+1<len(sub) &&a+1<len(M) {
                count++
                sum+=M[a+1][b+1]
            }
            res[a][b] = sum/count
        }
    }
    return res
}
