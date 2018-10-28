func generateParenthesis(n int) []string {
    res:=[]string{}
    search(n,n,"",&res)
    return res
}

func search(left,right int, str string,res *[]string){
    if left<0 || right<0||left>right {return}
    if left==0 && right==0{
        *res = append(*res,str)
        return
    }
    search(left-1,right,str+"(",res)
    search(left, right-1,str+")",res)
}
