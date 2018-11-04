func minSwapsCouples(row []int) int {
    res:=0
    for i:=0;i<len(row);i+=2{
        if row[i+1]==row[i]^1{ continue }
        for j:=i+2;j<len(row);j++{
            if row[j] == row[i]^1 {
                row[i+1], row[j] = row[j],row[i+1]
            }
        }
        res++
    }
    return res
}
