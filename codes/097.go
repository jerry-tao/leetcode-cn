func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2)!=len(s3){
		return false
	}

	var dfs func(c1, c2, c3 int) bool

	dfs = func(c1, c2, c3 int) bool {
		if c3 == len(s3) {
			return true
		}

		if (c1 == len(s1) && s3[c3] != s2[c2]) || (c2 == len(s2) && s3[c3] != s1[c1]) {
			return false
		}

		if (c1 < len(s1) && s3[c3] == s1[c1] && dfs(c1+1, c2, c3+1)) || (c2 < len(s2) && s3[c3] == s2[c2] && dfs(c1, c2+1, c3+1) ) {
			return true
		}

		return false
	}

	return dfs(0, 0, 0)
}
