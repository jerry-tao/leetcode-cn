func complexNumberMultiply(a string, b string) string {
	p1, p2 := 0, 0
	var sa []string
	var sb []string
	if strings.Contains(a, "+-") {
		sa = strings.Split(a, "+-");
		p1 = 1
	} else {
		sa = strings.Split(a, "+")
	}
	if strings.Contains(b, "+-") {
		sb = strings.Split(b, "+-");
		p2 = 1
	} else {
		sb = strings.Split(b, "+")
	}
	v1, _ := strconv.Atoi(sa[0])
	v2, _ := strconv.Atoi(strings.Split(sa[1], "i")[0])
	v3, _ := strconv.Atoi(sb[0])
	v4, _ := strconv.Atoi(strings.Split(sb[1], "i")[0])
	r1, r2 := 0, 0
	p3 := 0
	if p1 == p2 {
		r1 = v1*v3 - v2*v4
		if p1 == 0 {
			r2 = v1*v4 + v2*v3
		} else {
            r2 = -(v1*v4 + v2*v3)
		}
        if r2<0{
            r2 = -r2
            p3 = 1
        }
	} else {
		r1 = v1*v3 + v2*v4
		if p1 == 0 {
			r2 = v2*v3 - v1*v4
		} else {
			r2 = v1*v4 - v2*v3
		}
		if r2 < 0 {
            r2=-r2
			p3 = 1
		}
	}
	if p3 == 0 {
		return strconv.Itoa(r1) + "+" + strconv.Itoa(r2) + "i"
	} else {
		return strconv.Itoa(r1) + "+-" + strconv.Itoa(r2) + "i"
	}
}
