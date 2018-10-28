func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	dfs(&res, k, 1, n, []int{})
	return res
}
func dfs(res *[][]int, k int, number int, n int, out []int) {
	if len(out) == k && n == 0 {
		fmt.Println("here")
		o := make([]int, len(out))
		copy(o, out)
		*res = append((*res), o)
		return
	}
	for i := number; i <= 9; i++ {
		if n-i >= 0 {
			out = append(out, i)
			dfs(res, k, i+1, n-i, out)
            out = out[:len(out)-1]
            fmt.Println("here",out,i)
        }else{
            break
        }
	}

}
