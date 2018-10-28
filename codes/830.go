func largeGroupPositions(S string) [][]int {
	s, e := 0, 0
	res := [][]int{}
	for e < len(S) {
		if S[e] != S[s] {
			if e-s >= 3 {
				res = append(res, []int{s, e - 1})
			}
			s = e
		}
		e++
	}

    if e-s>=3{
        res = append(res,[]int{s,e-1})
    }
	return res
}
