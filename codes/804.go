func uniqueMorseRepresentations(words []string) int {
    mos:= []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    tmp:=""
    m:=map[string]int{}
    for _,v:=range words {
        tmp=""
        for _,c:=range v{
            tmp+=mos[int(c-97)]
        }
        m[tmp] = 1
    }
    return len(m)
}
