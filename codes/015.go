func threeSum(nums []int) [][]int {
    res:=[][]int{}
    sort.Ints(nums)
    n:=len(nums)
    for a:=0;a<len(nums);a++{
        if nums[a]>0{break}
        if a>0 && nums[a]==nums[a-1]{continue}
        l,r:=a+1,n-1
        for l<r{
            if nums[a]+nums[l]+nums[r]==0{
                res = append(res,[]int{nums[a],nums[l],nums[r]})
                for l<r && nums[l]==nums[l+1]{l++}
                for l<r && nums[r]==nums[r-1]{r--}
                l++
                r--
            }else if nums[a]+nums[l]+nums[r]<0{l++}else{r--}
        }
    }

    return res
}
