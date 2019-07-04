// 暴力破解还击败了72%
func findLongestWord(s string, d []string) string {
	var tmp []string
	maxLength := 0
	for _, v := range d {
		if len(v) > len(s) {
			break
		}
		j := 0
		valid := false
		for i := 0; i < len(v); i++ {
			for j < len(s) {
				if s[j] == v[i] {
					if i == len(v)-1 {
						valid = true
					}
					j++
					break
				}
				j++
			}
		}
		if valid {
			if len(v) > maxLength {
				tmp = tmp[:0]
				maxLength = len(v)
				tmp = append(tmp, v)
			} else if len(v) == maxLength {
				tmp = append(tmp, v)
			}
		}
	}

	if len(tmp) == 1 {
		return tmp[0]
	}
	if len(tmp) > 1 {
		sort.Strings(tmp)
		return tmp[0]
	}
	return ""
}
