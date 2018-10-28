func calPoints(ops []string) int {
    res:=0
    stack:=[]int{}
    for _,v:=range ops{
        if i,err:=strconv.Atoi(v);err==nil{
            stack = append(stack,i)
        }else{
            switch v{
                case "C":
                    stack = stack[:len(stack)-1]
                case "D":
                stack = append(stack,stack[len(stack)-1]*2)
                case "+":
                stack = append(stack,stack[len(stack)-1]+stack[len(stack)-2])
            }
        }

    }
    for _,v:=range stack{
        res+=v
    }
    return res
}
