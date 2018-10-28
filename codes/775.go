func isIdealPermutation(A []int) bool {
    for i,v:=range A{
        if v-i>1 || v-i<(-1){
            return false
        }
    }
    return true
}
