func lengthOfLastWord(s string) int {
    if len(s) == 0{
        return 0
    }
    str:=strings.Split(s," ")

    for k := len(str)-1;k>=0;k--{
        if len(str[k]) != 0{
            return len(str[k])
      }
    }
return 0
}
