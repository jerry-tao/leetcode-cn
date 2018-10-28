func sortColors(nums []int)  {
    l:=0
    r:=len(nums)-1
    if l==r{
        return
    }
    for i:=0;i<=r;i++{
        v:=nums[i]
        if v==0{
            swap(nums,i,l)
            l++
        }
        if v==2{
            swap(nums,i,r)
            r--
            i--
        }
    }
}

func swap(nums []int,a,b int){
    tmp:=nums[a]
    nums[a]=nums[b]
    nums[b]=tmp
}
