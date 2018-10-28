func longestMountain(A []int) int {
    max := 0
    tmp := 0
    for i:=1;i<len(A)-1;i++{
        if A[i]>A[i-1] && A[i]>A[i+1]{
            tmp = 1
            for j:=i-1;j>=0;j--{
                if A[j]<A[j+1]{
                    tmp++
                }else{
                    break
                }
            }
            for k:=i+1;k<len(A);k++{
                if A[k]<A[k-1]{
                    tmp++
                }else{
                    break
                }
            }
            if tmp>max{
                max=tmp
            }
        }
    }
    return max
}
