// on级别，但是on级别的额外空间
func topKFrequent(nums []int, k int) []int {
	re := make([][]int, len(nums)+1)
	for i := range re {
		re[i] = []int{}
	}
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums {
		m[v]++
	}
	for key, v := range m {
        re[v] = append(re[v],key)
	}
	i := 0
	for j := len(nums); j >= 0 && i < k; j-- {
		if len(re[j]) != 0 {
			res = append(res, re[j]...)
			i += len(re[j])
		}
	}
	return res
}
