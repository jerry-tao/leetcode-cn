func mostCommonWord(paragraph string, banned []string) string {
    m := make(map[string]int)

    for i:=range banned{ m[banned[i]]=-1 }
    max,res:=0,""
    f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
    paragraph = strings.ToLower(paragraph)
    words := strings.FieldsFunc(paragraph, f)
    for _,v:=range words {
        if m[v]<0{
            continue
        }
        m[v]+=1
        if m[v]>max{
            max = m[v]
            res=v
        }
	}
    return res
}
