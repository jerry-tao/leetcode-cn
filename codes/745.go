// 如果用字典树会有很大的内存优化空间，不过go里实现字典树太麻烦了
type WordFilter struct {
	pre map[string][]int
	suf map[string][]int
}

func Constructor(words []string) WordFilter {
	w := WordFilter{
		make(map[string][]int), make(map[string][]int),
	}
	for i, v := range words {
		tmp := ""
		w.pre[tmp] = append(w.pre[tmp], i)
		for _, r := range v {
			tmp += string(r)
			w.pre[tmp] = append(w.pre[tmp], i)
		}

		tmp = ""
		w.suf[tmp] = append(w.suf[tmp], i)
		for j := len(v) - 1; j >= 0; j-- {
			tmp = string(v[j]) + tmp
			w.suf[tmp] = append(w.suf[tmp], i)
		}
	}
	return w
}

func (this *WordFilter) F(prefix string, suffix string) int {
	v1, ok1 := this.pre[prefix]
	v2, ok2 := this.suf[suffix]
	if ok1 && ok2 {
		for i := len(v1) - 1; i >= 0; i-- {
			for j := len(v2) - 1; j >= 0; j-- {
				if v1[i] == v2[j] {
					return v1[i]
				}
			}
		}
	}
	return -1
}
