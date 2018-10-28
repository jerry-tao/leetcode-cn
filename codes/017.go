func letterCombinations(digits string) []string {
    if len(digits)==0{return nil}
    res:=[]string{}
    dict:=[]string {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

    var dfs func(int,string)
    dfs = func(s int,o string){
        if s==len(digits){
            res = append(res, o)
            return
        }
        str := dict[digits[s] - '0']
        for j:=0;j<len(str);j++{
            dfs(s+1, o + string(str[j]))
        }
    }
    dfs(0,"")
    return res
}
