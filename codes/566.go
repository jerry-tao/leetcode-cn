func matrixReshape(nums [][]int, r int, c int) [][]int {
    n:=len(nums)
    m:=len(nums[0])
    if n*m!=r*c{
        return nums
    }
    q,a:=0,0
    res:=make([][]int, r)
    for i:=0;i<r;i++{
        res[i] = make([]int,c)
    }
    for _,v:=range nums{
        for _,v2:=range v{
            res[q][a] = v2
            a++
            if a==c{
                a=0
                q++
            }
        }
    }
    return res
}
