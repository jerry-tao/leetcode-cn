func maxArea(height []int) int {
    l:=0
    r:=len(height)-1
    res:=0
    for l<r{
        if (r-l)*min(height[r],height[l])>res{
            res = (r-l)*min(height[r],height[l])
        }
        if height[r]<height[l]{
            r--
        }else{
            l++
        }
    }
    return res
}
func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}
