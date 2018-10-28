func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
    for i:=0;i<len(numbers);i++{
        if r,ok:=m[numbers[i]];ok{
            return []int{r+1,i+1}
        }else{
            m[target-numbers[i]] = i
        }
    }
    return nil
}
