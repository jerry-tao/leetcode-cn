func toGoatLatin(S string) string {
    str := strings.Split(S, " ")
    for i,_:= range str{
        if strings.IndexByte("aeiouAEIOU",str[i][0])>=0{
            str[i]+="ma"
        }else{
            str[i] = str[i][1:]+string(str[i][0])+"ma"
        }

        str[i]+=strings.Repeat("a", i+1)
    }
    return strings.Join(str," ")
}
