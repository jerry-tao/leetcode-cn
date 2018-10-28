func combine(n int, k int) [][]int {
	res:=[][]int{}
	out:=[]int{}
	var dfs func(int)
	dfs =func(lc int){
		if len(out)==k{
			res = append(res,c(out))
			return
		}
		for i:=lc;i<=n;i++{
			out=append(out,i)
			dfs(i+1)
			out=out[:len(out)-1]
		}
	}
	dfs(1)
	return res
}
func c(i []int)[]int{
	c:=make([]int,len(i))
	copy(c,i)
	return c
}
