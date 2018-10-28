func nextGreatestLetter(letters []byte, target byte) byte {
    if target>=letters[len(letters)-1]{return letters[0]}
    for _,v:=range letters{
        if target<v{
            return v
        }
    }
    return letters[0]
}
