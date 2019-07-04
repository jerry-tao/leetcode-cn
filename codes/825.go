func numFriendRequests(ages []int) int {
	m := make([]int, 121)
	r := 0
	for _, i := range ages {
		m[i] += 1
	}
	for i, p := range m {
		if p > 1 && i>=i/2+8 {
			r += p * (p - 1)
		}

		for j := i/2 + 8; j < i; j++ {
			r += p * m[j]
		}
	}
	return r
}
