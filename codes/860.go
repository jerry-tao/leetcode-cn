func lemonadeChange(bills []int) bool {
    res:=[]int{}
    for _,v:=range bills{
        if v-5>0{
            if len(res)==0{return false}
            t:=v-5
            for j:=len(res)-1;j>=0;j--{
                if t>=res[j]{
                    t=t-res[j]
                    res[j]=0
                }
            }
            if t!=0{return false}
            res = append(res,v)
        }else{

            res = append(res, v)
            sort.Ints(res)
        }
    }
    return true
}
