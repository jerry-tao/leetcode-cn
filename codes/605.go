func canPlaceFlowers(flowerbed []int, n int) bool {
    if n==0{return true}
    l := len(flowerbed)
    if flowerbed[0]==0{ flowerbed = append([]int{0},flowerbed...)}
    if flowerbed[l-1]==0{ flowerbed = append(flowerbed,0)}
    if len(flowerbed)<3{return false}
    for i:=1;i<len(flowerbed)-1;i++{
        if flowerbed[i]+flowerbed[i-1]+flowerbed[i+1] == 0{
            n--
            if n==0{return true}
            flowerbed[i]=1
        }
    }
    return n<=0
}
