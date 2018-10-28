func merge(nums1 []int, m int, nums2 []int, n int)  {
    t:=make([]int,m+n)
    a,b:=0,0
    for i:=0;i<m+n;i++{
        if a<m&&b<n{
            if nums1[a]<=nums2[b]{
                t[i] = nums1[a]
                a++
            }else{
                t[i] = nums2[b]
                b++
            }
        }else if a<m && b>=n{
            t[i] = nums1[a]
            a++
        }else if a>=m && b<n{
            t[i] = nums2[b]
            b++
        }
    }
    for i,v:=range t{
        nums1[i] = v
    }
}
